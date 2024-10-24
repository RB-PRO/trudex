package bitrix_consumer

import (
	"context"
	correlationid "github.com/dmytrohridin/correlation-id"
	"github.com/gin-gonic/gin"
)

// middleware add correlation id
func correlationID(c *gin.Context) {
	if id := correlationid.FromContext(c.Request.Context()); id == "" {
		c.Request.WithContext(
			context.WithValue(c.Request.Context(),
				correlationid.ContextKey, correlationid.New(),
			),
		)
	}
}
