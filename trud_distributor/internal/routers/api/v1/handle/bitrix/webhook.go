package bitrix

import (
	"github.com/gin-gonic/gin"
	"io"
	"trudex/common/config"
	"trudex/trud_distributor/internal"
	"trudex/trud_distributor/internal/services/dto"
	"trudex/trud_distributor/internal/services/rabbitmq"
)

func HandleBitrixConsumer(service *rabbitmq.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			dto.NewInternalErr(c, err, "error read body '%s'", string(body))
			return
		}

		ctx, _, err := config.LoadToCtxFromKeys[internal.Config](c.Request.Context(), c.Keys)
		if err != nil {
			dto.NewInternalErr(c, err, "error load config")
			return
		}

		//// TODO: Add parser query
		//var data dto.Root
		//m, err := url.ParseQuery(string(body))
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//_ = m
		//if err := json.Unmarshal(body, &data); err != nil {
		//	dto.NewInternalErr(c, err, "error unmarshal object body data '%s'", string(body))
		//	return
		//}

		// service > put in rabbit mq
		isSend, err := service.Push(ctx)
		if err != nil {
			// todo: handle error me
			dto.NewInternalErr(c, err, "internal rabbit mq error")
			return
		}

		if !isSend {
			dto.NewInternalErr(c, err, "internal rabbit mq error")
			return
		}

		dto.NewCompleteResponse(c)
		return
	}
}
