package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	appconfig "trudex/trud_distributor/internal"
)

const (
	ConfigCtxKey = "cfg"
)

// Config is middleware for add config id in context
func Config(ctx context.Context) gin.HandlerFunc {

	cfg, err := appconfig.LoadConfigFromContext(ctx)
	if err != nil {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		c.Set(ConfigCtxKey, cfg)
		c.Next()
	}
}
