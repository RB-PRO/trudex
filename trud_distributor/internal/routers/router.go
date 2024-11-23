package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trudex/common/logger"
	bitrixHandle "trudex/trud_distributor/internal/routers/api/v1/handle/bitrix"
	"trudex/trud_distributor/internal/routers/api/v1/middleware"
	"trudex/trud_distributor/internal/services"
)

func InitRouter(log *logger.Logger, services services.Services) *http.Server {
	//r := gin.New()
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger(log))
	r.Use(middleware.CorrelationID())

	r.POST("/", bitrixHandle.HandleBitrixConsumer(services.RabbitmqService))

	// Create a server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return srv
}
