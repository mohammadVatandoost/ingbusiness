package main

import (
	"context"
	"fmt"
	"sync"
	"syscall"
	"time"

	"git.cafebazaar.ir/divar/cloud-sand-boxing/internal/database"
	"git.cafebazaar.ir/divar/cloud-sand-boxing/internal/experiment"
	"git.cafebazaar.ir/divar/cloud-sand-boxing/internal/serviceroute"
	"git.cafebazaar.ir/divar/cloud-sand-boxing/pkg/logger"

	"git.cafebazaar.ir/divar/cloud-sand-boxing/internal/clients"

	"git.cafebazaar.ir/divar/cloud-sand-boxing/internal/app/controller"
	grpcAPI "git.cafebazaar.ir/divar/cloud-sand-boxing/internal/core/grpc"
	cntext "git.cafebazaar.ir/divar/cloud-sand-boxing/pkg/context"
	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	_ "github.com/GoAdminGroup/go-admin/adapter/echo"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	_ "github.com/GoAdminGroup/themes/adminlte"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serve(cmd, args); err != nil {
			logrus.WithError(err).Fatal("Failed to serve.")
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) error {

	printVersion()

	conf := loadConfigOrPanic(cmd)
	configureLoggerOrPanic(conf.Logger)

	log := logger.NewLogger()

	//prometheusMetricServer := prometheus.StartMetricServerOrPanic(conf.Metric.ListenPort)
	//defer prometheus.ShutdownMetricServerOrPanic(prometheusMetricServer)

	db, err := database.NewDBConnection(log, conf.Postgres, uint(conf.Postgres.MigrationVersion))
	if err != nil {
		return fmt.Errorf("failed to create DB connection: %v", err.Error())
	}

	cFlag, err := clients.NewControlPlaneFlag(time.Duration(conf.GRPC.TimeOut) * time.Second)
	if err != nil {
		logrus.WithError(err).Panic("Failed to NewControlPlaneFlag")
	}

	cOverride, err := clients.NewControlPlaneOverride(time.Duration(conf.GRPC.TimeOut) * time.Second)
	if err != nil {
		logrus.WithError(err).Panic("Failed to NewControlPlaneOverride")
	}

	cCondition, err := clients.NewControlPlaneCondition(time.Duration(conf.GRPC.TimeOut) * time.Second)
	if err != nil {
		logrus.WithError(err).Panic("Failed to NewControlPlaneCondition")
	}

	cService, err := clients.NewControlPlaneService(time.Duration(conf.GRPC.TimeOut) * time.Second)
	if err != nil {
		logrus.WithError(err).Panic("Failed to NewControlPlaneService")
	}

	experimentDirectory := experiment.NewDirectory(log, db)
	serviceRouteDirectory := serviceroute.NewDirectory(log, db)

	controlPlane := controller.New(experimentDirectory, serviceRouteDirectory,
		cFlag, cCondition, cOverride, cService)

	grpcService := grpcAPI.New(experimentDirectory, serviceRouteDirectory, controlPlane, conf.GRPC)
	grpcServer := getGrpcServer(grpcService, []grpc.UnaryServerInterceptor{})

	serverContext, serverCancel := cntext.WithSignalCancellation(
		context.Background(),
		syscall.SIGTERM, syscall.SIGINT,
	)
	defer serverCancel()

	var serverWaitGroup sync.WaitGroup
	serverWaitGroup.Add(1)

	go func() {
		defer serverWaitGroup.Done()
		startGrpcServerOrPanic(conf.GRPC.ListenPort, grpcServer)
	}()

	<-serverContext.Done()
	go func() {
		grpcServer.GracefulStop()
	}()

	serverWaitGroup.Wait()
	return nil

}
