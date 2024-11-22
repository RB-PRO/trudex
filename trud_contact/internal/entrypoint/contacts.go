package entrypoint

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"trudex/trud_contact/internal/contact"
)

type Service struct {
	*logrus.Logger
	*route.Service
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
	consumer := route.NewService(8081, logger, contact.HandleFunction)

	return &Service{
		Logger:  logger,
		Service: consumer,
	}
}

func (s *Service) Run() {
	s.RunConsumer()
}
