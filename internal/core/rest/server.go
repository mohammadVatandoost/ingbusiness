package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/authentication"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/frequentmessages"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/ingmessenger"
	"github.com/mohammadVatandoost/ingbusiness/pkg/logger"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Engine                  *gin.Engine
	authenticationService   *authentication.Service
	messengerService        *ingmessenger.Service
	frequentMessagesService *frequentmessages.Service
	logger                  *logrus.Logger
}

func New(authenticationService *authentication.Service,
	messengerService *ingmessenger.Service, frequentMessagesService *frequentmessages.Service) *Server {
	return &Server{
		authenticationService:   authenticationService,
		messengerService:        messengerService,
		frequentMessagesService: frequentMessagesService,
		logger:                  logger.NewLogger(),
		Engine:                  gin.New(),
	}
}
