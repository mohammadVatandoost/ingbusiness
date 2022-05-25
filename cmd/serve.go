package main

import (
	"context"
	"fmt"
	"github.com/mohammadVatandoost/ingbusiness/internal/access"
	restAPI "github.com/mohammadVatandoost/ingbusiness/internal/core/rest"
	"github.com/mohammadVatandoost/ingbusiness/internal/database"
	"github.com/mohammadVatandoost/ingbusiness/internal/ingaccounts"
	"github.com/mohammadVatandoost/ingbusiness/internal/organization"
	roles "github.com/mohammadVatandoost/ingbusiness/internal/role"
	"github.com/mohammadVatandoost/ingbusiness/internal/savedmessages"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/authentication"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/frequentmessages"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/iam"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/ingmessenger"
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
	"github.com/mohammadVatandoost/ingbusiness/pkg/logger"
	"sync"
	"syscall"
	"time"

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
	savedMessagesDirectory := savedmessages.NewDirectory(log, db)
	ingAccountsDirectory := ingaccounts.NewDirectory(log, db)
	organizationDirectory := organization.NewDirectory(log, db)
	accessDirectory := access.NewDirectory(log, db)
	rolesDirectory := roles.NewDirectory(log, db)

	authenticationService := authentication.New(log, usersDirectory, conf.Auth)
	ingMessengerService := ingmessenger.New(usersDirectory, ingAccountsDirectory, savedMessagesDirectory)
	frequentMessagesService := frequentmessages.New(savedMessagesDirectory)
	iamService := iam.New(organizationDirectory, rolesDirectory, accessDirectory, usersDirectory)

	serverREST := restAPI.New(log, conf.Rest, authenticationService, iamService, ingMessengerService, frequentMessagesService)
	serverREST.Routes()

	serverContext, serverCancel := cntext.WithSignalCancellation(
		context.Background(),
		syscall.SIGTERM, syscall.SIGINT,
	)
	defer serverCancel()

	var serverWaitGroup sync.WaitGroup
	serverWaitGroup.Add(1)

	go func() {
		defer serverWaitGroup.Done()
		serverREST.Run()
		//startGrpcServerOrPanic(conf.GRPC.ListenPort, grpcServer)
	}()

	<-serverContext.Done()
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		serverREST.Shutdown(ctx)
	}()

	serverWaitGroup.Wait()
	return nil

}
