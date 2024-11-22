package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"trudex/trud_distributor/internal/routers/api/v1/middleware"
)

func NewEngine(
	port int,
	log *logrus.Logger,
) *Service {
	r := gin.Default()
	r.Use(middleware.CorrelationID())

	service := &Service{
		Engine: r,
		Logger: log,
		port:   port,
	}
	r.POST("/", service.handleProducer)
	return service
}
