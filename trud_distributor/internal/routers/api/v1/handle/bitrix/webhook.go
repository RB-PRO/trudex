package bitrix

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"trudex/trud_distributor/internal/services/dto"
	"trudex/trud_distributor/internal/services/rabbitmq"
	"trudex/trud_distributor/pkg/body_parser"
)

func HandleBitrixConsumer(service *rabbitmq.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		bodyByte, err := io.ReadAll(c.Request.Body)
		if err != nil {
			dto.NewInternalErr(c, err, "error read body '%s'", string(bodyByte))
			return
		}

		event, err := body_parser.ParseEventData(bodyByte)
		if err != nil {
			dto.NewInternalErr(c, err, "error parse body '%s'", string(bodyByte))
		}

		fmt.Printf("$+$ %+v\n", string(bodyByte))
		fmt.Printf("$+$ %+v\n", event)

		// service > put in rabbit mq
		if err := service.Push(ctx, event); err != nil {
			// todo: handle error me
			logrus.WithContext(ctx).Error(err)
			dto.NewInternalErr(c, err, "internal rabbit mq error")
			return
		}

		dto.NewCompleteResponse(c)
		return
	}
}
