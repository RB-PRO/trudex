package bitrix

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/url"
	"trudex/common/config"
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

		z := config.LoadFromContext(c.Request.Context())
		_ = z

		fmt.Println("$+$")

		// TODO: Add parser query
		var data dto.Root
		m, err := url.ParseQuery(string(body))
		if err != nil {
			log.Fatal(err)
		}

		_ = m
		if err := json.Unmarshal(body, &data); err != nil {
			dto.NewInternalErr(c, err, "error unmarshal object body data '%s'", string(body))
			return
		}

		// service > put in rabbit mq
		if ok, err := service.Push(c.Request.Context()); err != nil || !ok {
			dto.NewInternalErr(c, err, "internal rabbit mq error")
			return
		}

		dto.NewCompleteResponse(c)
		return
	}
}
