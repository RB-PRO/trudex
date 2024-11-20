package bitrix_producer

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
)

type BitrixWebhookData struct {
	Event                string `url:"event"`
	EventHandlerID       string `url:"event_handler_id"`
	DataFieldsID         string `url:"data[FIELDS][ID]"`
	DataFieldsTypeID     string `url:"data[FIELDS][ENTITY_TYPE_ID]"`
	Timestamp            int64  `url:"ts"`
	AuthDomain           string `url:"auth[domain]"`
	AuthClientEndpoint   string `url:"auth[client_endpoint]"`
	AuthServerEndpoint   string `url:"auth[server_endpoint]"`
	AuthMemberID         string `url:"auth[member_id]"`
	AuthApplicationToken string `url:"auth[application_token]"`
}

func (s *Service) handleProducer(c *gin.Context) {
	ctx := c.Request.Context()
	s.WithContext(ctx).Infof("request received in patch '%s'", c.FullPath())

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		NewInternalErr(c, err, "error read body request")
		return
	}

	var data BitrixWebhookData
	if err := json.Unmarshal(body, &data); err != nil {
		NewInternalErr(c, err, "error unmarshal object body data '%s'", string(body))
		return
	}

	// service > put in rabbit mq

	NewCompleteResponse(c)
	return
}
