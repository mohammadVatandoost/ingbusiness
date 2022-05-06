package rest

import (
	"git.cafebazaar.ir/divar/cloud-sand-boxing/internal/app/controller"
	"git.cafebazaar.ir/divar/cloud-sand-boxing/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Engine       *gin.Engine
	ControlPlane *controller.ControlPlane
	logger       *logrus.Logger
}

func New(ControlPlane *controller.ControlPlane) *Server {
	return &Server{
		ControlPlane: ControlPlane,
		logger:       logger.NewLogger(),
		Engine:       gin.New(),
	}
}
