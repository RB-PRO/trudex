package bitrix_producer

import (
	"fmt"
	corr_id "github.com/dmytrohridin/correlation-id"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Status    int    `json:"status,omitempty"`
	RequestID string `json:"request_id,omitempty"`
	Message   string `json:"message,omitempty"`
	Error     error  `json:"error,omitempty"`
}

func newResponse(c *gin.Context, status int, err error, formatError string, args ...string) response {
	return response{
		Status:    status,
		RequestID: corr_id.FromContext(c.Request.Context()),
		Message:   fmt.Sprintf(formatError, args),
		Error:     err,
	}
}

func NewCompleteResponse(c *gin.Context) {
	c.JSON(http.StatusOK,
		newResponse(c, http.StatusOK, nil, ""),
	)
}

func NewInternalErr(c *gin.Context, err error, format string, args ...string) {
	c.JSON(http.StatusInternalServerError,
		newResponse(c, http.StatusInternalServerError, err, format, args...),
	)
}
