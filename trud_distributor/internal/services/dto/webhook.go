package dto

type Root struct {
	Event          string `json:"event"`
	EventHandlerID string `json:"event_handler_id"`
	Data           Data   `json:"data"`
	TS             string `json:"ts"`
	Auth           Auth   `json:"auth"`
}
type Auth struct {
	Domain           string `json:"domain"`
	ClientEndpoint   string `json:"client_endpoint"`
	ServerEndpoint   string `json:"server_endpoint"`
	MemberID         string `json:"member_id"`
	ApplicationToken string `json:"application_token"`
}

type Data struct {
	Fields map[string]string `json:"FIELDS"`
}
