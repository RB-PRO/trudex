package bitrix_producer

import (
	"context"
	corr_id "github.com/dmytrohridin/correlation-id"
	"github.com/gin-gonic/gin"
)

// middleware add correlation id
func correlationID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if id := corr_id.FromContext(c.Request.Context()); id == "" {
			c.Request.WithContext(
				context.WithValue(c.Request.Context(),
					corr_id.ContextKey, corr_id.New(),
				),
			)
		}
		c.Next()
	}
}
