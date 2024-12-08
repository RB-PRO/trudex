package internal

import (
	"trudex/trud_distributor/internal/services/rabbitmq"
)

type Config struct {
	Route        RouteConfig           `yaml:"route"`
	RabbitConfig rabbitmq.RabbitConfig `yaml:"rabbit_config"`
}

type RouteConfig struct {
	Port int `yaml:"port"`
}
