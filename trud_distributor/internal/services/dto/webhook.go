package dto

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
