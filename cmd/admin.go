package main

import (
	"context"
	"fmt"
	"github.com/mohammadVatandoost/ingbusiness/internal/database"
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
	"github.com/mohammadVatandoost/ingbusiness/pkg/logger"
	"sync"
	"syscall"

	grpcAPI "github.com/mohammadVatandoost/ingbusiness/internal/core/grpc"
	restAPI "github.com/mohammadVatandoost/ingbusiness/internal/core/rest"
	"github.com/mohammadVatandoost/ingbusiness/internal/goadmin"
	cntext "github.com/mohammadVatandoost/ingbusiness/pkg/context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "start admin server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serveAdmin(cmd, args); err != nil {
			logrus.WithError(err).Fatal("Failed to serve.")
		}
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
}

func serveAdmin(cmd *cobra.Command, args []string) error {
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
	usersDirectory := users.NewDirectory(log, db)

	//controlPlane := controller.New(experimentDirectory, serviceRouteDirectory, cFlag, cCondition, cOverride, cService)

	serverREST := restAPI.New(controlPlane)
	serverREST.Routes()

	goAdmin := goadmin.NewController(experimentDirectory, serviceRouteDirectory, controlPlane)
	go func() {
		e := goAdmin.ServeAdmin(conf.GoAdmin, database.LoadConfig(), serverREST.Engine)
		if e != nil {
			logrus.Errorf("go admin can not serve, err: %v", e.Error())
		}
	}()

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
