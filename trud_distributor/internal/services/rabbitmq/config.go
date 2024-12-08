package rabbitmq

// RabbitConfig for establishing one connection
type RabbitConfig struct {
	Name      string     `yaml:"name"`
	Consumers []Consumer `yaml:"consumers"`
}

type Consumer struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}
