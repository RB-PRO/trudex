package routers

import (
	"github.com/gin-gonic/gin"
	"trudex/common/logger"
	bitrixHandle "trudex/trud_distributor/internal/routers/api/v1/handle/bitrix"
	"trudex/trud_distributor/internal/routers/api/v1/middleware"
	"trudex/trud_distributor/internal/services"
)

func InitRouter(log *logger.Logger, services services.Services) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger(log))
	r.Use(middleware.CorrelationID())

	r.POST("/", bitrixHandle.HandleBitrixConsumer(services.RabbitmqService))

	return r
}
