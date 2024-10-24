package bitrix_consumer

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Service struct {
	*gin.Engine
	*logrus.Logger
	handlerWorkFunc func(ctx context.Context) (string, error)
	port            int
}

func NewService(
	port int,
	log *logrus.Logger,
	handlerWorkFunc func(ctx context.Context) (string, error),
) *Service {
	r := gin.Default()
	r.Use(correlationID)

	service := &Service{
		Engine:          r,
		Logger:          log,
		handlerWorkFunc: handlerWorkFunc,
		port:            port,
	}

	r.POST("/", service.handleConsumer)
	return service
}

func (s *Service) RunConsumer() {
	_ = s.Run(fmt.Sprintf(":%d", s.port))
}
