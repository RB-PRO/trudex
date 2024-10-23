package bitrix_consumer

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	*gin.Engine
	*logrus.Logger
}

func NewService(
	log *logrus.Logger,
	handlerFunc func(ctx context.Context) (string, error),
) *Service {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		ctx := c.Request.Context()
		log.WithContext(ctx).Infof("request received in patch '%s'", c.FullPath())

		result, err := handlerFunc(ctx)
		if err != nil {
			fmt.Print("$+$ err")
		}
		fmt.Print("$+$", result)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})
	r.POST("/", func(c *gin.Context) {
		ctx := c.Request.Context()
		log.WithContext(ctx).Infof("request received in patch '%s'", c.FullPath())

		result, err := handlerFunc(ctx)
		if err != nil {
			fmt.Print("$+$ err")
		}
		fmt.Print("$+$", result)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})
	return &Service{Engine: r, Logger: log}
}

func (s *Service) RunConsumer() {
	s.Info("start bitrix producer")
	s.Run(":8081")
}
