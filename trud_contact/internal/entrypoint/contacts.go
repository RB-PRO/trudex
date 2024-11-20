package entrypoint

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"trudex/common/bitrix_producer"
	"trudex/trud_contact/internal/contact"
)

type Service struct {
	*logrus.Logger
	*bitrix_producer.Service
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
	consumer := bitrix_producer.NewService(8081, logger, contact.HandleFunction)

	return &Service{
		Logger:  logger,
		Service: consumer,
	}
}

func (s *Service) Run() {
	s.RunConsumer()
}
