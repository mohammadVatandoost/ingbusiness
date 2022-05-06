package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammadVatandoost/ingbusiness/pkg/logger"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Engine *gin.Engine
	logger *logrus.Logger
}

func New() *Server {
	return &Server{
		logger: logger.NewLogger(),
		Engine: gin.New(),
	}
}
