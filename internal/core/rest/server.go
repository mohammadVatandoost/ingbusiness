package rest

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/authentication"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/frequentmessages"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/iam"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/ingmessenger"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	Engine                  *gin.Engine
	authenticationService   *authentication.Service
	messengerService        *ingmessenger.Service
	frequentMessagesService *frequentmessages.Service
	iamService              *iam.Service
	logger                  *logrus.Logger
	srv                     *http.Server
	conf                    Config
}

func (s *Server) Run() {
	port := s.conf.ListenPort
	addr := fmt.Sprintf(":%v", port)
	s.logger.Infof("CRM REST API Service Running, addr: %v \n", addr)

	s.srv = &http.Server{
		Addr:    addr,
		Handler: s.Engine,
	}

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Panicf("error on listen HTTP Server, error: %s", err.Error())
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.Panicf("server forced to shutdown, err: %s", err.Error())
	}
}

func New(logger *logrus.Logger, authenticationService *authentication.Service, iamService *iam.Service,
	messengerService *ingmessenger.Service, frequentMessagesService *frequentmessages.Service) *Server {
	return &Server{
		authenticationService:   authenticationService,
		messengerService:        messengerService,
		frequentMessagesService: frequentMessagesService,
		iamService:              iamService,
		logger:                  logger,
		Engine:                  gin.New(),
	}
}
