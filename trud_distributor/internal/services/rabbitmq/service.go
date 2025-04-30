package rabbitmq

import (
	"context"
	"github.com/wagslane/go-rabbitmq"
	"log"
)

type Service struct {
	conn      *rabbitmq.Conn
	publisher *rabbitmq.Publisher
}

func NewService(cfg RabbitConfig) (*Service, func(ctx context.Context) error, error) {

	conn, err := rabbitmq.NewConn(
		"amqp://devops:Pozo2013!@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("trudex"),
	)
	if err != nil {
		log.Fatal(err)
	}

	stopFunc := func(ctx context.Context) error {
		// todo: implement me
		_ = conn.Close()
		publisher.Close()
		return nil
	}
	return &Service{
		conn:      conn,
		publisher: publisher,
	}, stopFunc, nil
}
