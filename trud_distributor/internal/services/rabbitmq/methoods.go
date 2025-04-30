package rabbitmq

import (
	"context"
	"github.com/wagslane/go-rabbitmq"
	"trudex/common/config"
)

func (s *Service) Push(ctx context.Context, params ...interface{}) error {
	cfg := config.LoadFromCtx[RabbitConfig](ctx)
	_ = cfg

	// get event

	return s.publisher.Publish(
		[]byte("hello, world"),
		[]string{"service-1-key"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("trudex"),
	)
}
