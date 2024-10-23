package entrypoint

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"trudex/common/bitrix_consumer"
	"trudex/trud_contact/internal/contact"
)

type Service struct {
	*logrus.Logger
	*bitrix_consumer.Service
}

func NewService() *Service {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	// init logger
	logger := logrus.New()

	// init consumer
	consumer := bitrix_consumer.NewService(logger, contact.HandleFunction)

	return &Service{
		Logger:  logger,
		Service: consumer,
	}
}

func (s *Service) Run() {
	s.RunConsumer()
}
