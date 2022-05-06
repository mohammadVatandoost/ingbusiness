package main

import (
	"context"
	"fmt"
	"github.com/mohammadVatandoost/ingbusiness/internal/database"
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
	"github.com/mohammadVatandoost/ingbusiness/pkg/logger"
	"sync"
	"syscall"

	cntext "github.com/mohammadVatandoost/ingbusiness/pkg/context"
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

	usersDirectory := users.NewDirectory(log, db)

	//controlPlane := controller.New(experimentDirectory, serviceRouteDirectory,
	//	cFlag, cCondition, cOverride, cService)

	//grpcService := grpcAPI.New(experimentDirectory, serviceRouteDirectory, controlPlane, conf.GRPC)
	//grpcServer := getGrpcServer(grpcService, []grpc.UnaryServerInterceptor{})

	serverContext, serverCancel := cntext.WithSignalCancellation(
		context.Background(),
		syscall.SIGTERM, syscall.SIGINT,
	)
	defer serverCancel()

	var serverWaitGroup sync.WaitGroup
	serverWaitGroup.Add(1)

	//go func() {
	//	defer serverWaitGroup.Done()
	//	startGrpcServerOrPanic(conf.GRPC.ListenPort, grpcServer)
	//}()

	<-serverContext.Done()
	//go func() {
	//	grpcServer.GracefulStop()
	//}()

	serverWaitGroup.Wait()
	return nil

}
