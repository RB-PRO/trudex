package bitrix_consumer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (s *Service) handleConsumer(c *gin.Context) {
	ctx := c.Request.Context()

	s.WithContext(ctx).Infof("request received in patch '%s'", c.FullPath())

	if jsonData, err := io.ReadAll(c.Request.Body); err == nil {
		fmt.Print("$+$  ", string(jsonData))
	}

	result, err := s.handlerWorkFunc(ctx)
	if err != nil {
		fmt.Print("$+$ err")
	}
	fmt.Print("$+$", result)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}
