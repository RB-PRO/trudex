package middleware

import (
	"github.com/gin-gonic/gin"
	"trudex/common/config"
	appconfig "trudex/trud_distributor/internal"
)

// Config is middleware for add config id in context and keys
func Config(configService *config.Service[appconfig.Config]) gin.HandlerFunc {
	cfg := configService.Config()

	return func(c *gin.Context) {
		c.Set(config.CtxKey, cfg)
		c.Next()
	}
}
