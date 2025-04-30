package rabbitmq

// RabbitConfig for establishing one connection
type RabbitConfig struct {
	ExchangeName     string     `yaml:"exchange_name"`
	ApplicationToken string     `yaml:"application_token"`
	Consumers        []Consumer `yaml:"consumers"`
}

type Consumer struct {
	Name       string `yaml:"name"`
	RoutingKey string `yaml:"routing_key"`
}
