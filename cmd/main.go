package main

import (
	"fmt"
	"net"
	"os"

	"github.com/mohammadVatandoost/ingbusiness/internal/config"
	"github.com/mohammadVatandoost/ingbusiness/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const serviceName = "instagram_helper"

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func loadConfigOrPanic(cmd *cobra.Command) *config.Config {
	conf, err := config.LoadConfig(cmd)
	if err != nil {
		logrus.WithError(err).Panic("Failed to load configurations")
	}
	return conf
}

func configureLoggerOrPanic(loggerConfig logger.Config) {
	if err := logger.Initialize(&loggerConfig); err != nil {
		logrus.WithError(err).Panic("Failed to configure logger")
	}
}

//func getGrpcServer(grpcServicer *grpcAPI.ServiceImplementation,
//	interceptors []grpc.UnaryServerInterceptor) *grpc.Server {
//
//	baseServer := server.NewGrpc(serviceName, server.WithInterceptor(interceptors...))
//	reflection.Register(baseServer)
//
//	return baseServer
//}

func startGrpcServerOrPanic(listenPort int, grpcServer *grpc.Server) {
	grpcAddr := fmt.Sprintf(":%d", listenPort)
	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logrus.WithError(err).Panic("Failed to listen")
		return
	}

	err = grpcServer.Serve(grpcListener)
	if err != nil {
		logrus.WithError(err).Panic("Failed to serve")
	}
}
