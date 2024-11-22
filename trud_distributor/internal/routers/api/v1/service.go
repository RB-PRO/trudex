package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"trudex/trud_distributor/internal/api/middleware"
)

type Service struct {
	*gin.Engine
	*logrus.Logger
	port int
}

func NewService(
	port int,
	log *logrus.Logger,
) *Service {
	r := gin.Default()
	r.Use(middleware.correlationID())

	service := &Service{
		Engine: r,
		Logger: log,
		port:   port,
	}
	r.POST("/", service.handleProducer)
	return service
}

func (s *Service) RunConsumer() {
	_ = s.Run(fmt.Sprintf(":%d", s.port))
}
