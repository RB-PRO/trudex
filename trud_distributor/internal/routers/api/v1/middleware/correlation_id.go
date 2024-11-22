package middleware

import (
	"context"
	correlationid "github.com/dmytrohridin/correlation-id"
	"github.com/gin-gonic/gin"
)

// CorrelationID is middleware for add correlation id in context
func CorrelationID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if id := correlationid.FromContext(c.Request.Context()); id == "" {
			c.Request.WithContext(
				context.WithValue(c.Request.Context(),
					correlationid.ContextKey, correlationid.New(),
				),
			)
		}
		c.Next()
	}
}
