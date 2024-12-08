package services

import (
	"trudex/common/config"
	"trudex/trud_distributor/internal"
	"trudex/trud_distributor/internal/services/rabbitmq"
)

type Services struct {
	ConfigService   *config.Service[internal.Config]
	RabbitmqService *rabbitmq.Service
}

func NewServices(
	configService *config.Service[internal.Config],
	rabbitmqService *rabbitmq.Service,
) Services {
	return Services{
		ConfigService:   configService,
		RabbitmqService: rabbitmqService,
	}
}
