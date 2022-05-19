package rest

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetDirectMessages(c *gin.Context) {
	experimentID := c.Param("id")
	id, err := strconv.ParseInt(experimentID, 10, 32)
	if err != nil {
		ErrorResponse(c, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

}

func (s *Server) SendDirectMessage(c *gin.Context) {
	experimentID := c.Param("id")
	id, err := strconv.ParseInt(experimentID, 10, 32)
	if err != nil {
		ErrorResponse(c, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//c.String(http.StatusOK, "Experiment %s is Enabled", experimentID)
}
