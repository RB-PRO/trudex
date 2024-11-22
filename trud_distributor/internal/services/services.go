package services

import "trudex/trud_distributor/internal/services/rabbitmq"

type Services struct {
	RabbitmqService *rabbitmq.Service
}

func NewServices(
	rabbitmqService *rabbitmq.Service,
) Services {
	return Services{
		RabbitmqService: rabbitmqService,
	}
}
