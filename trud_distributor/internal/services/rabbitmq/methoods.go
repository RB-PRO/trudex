package rabbitmq

import (
	"context"
	"trudex/common/config"
)

func (s *Service) Push(ctx context.Context, params ...interface{}) (bool, error) {
	cfg := config.LoadFromCtx[RabbitConfig](ctx)
	_ = cfg

	return false, nil
}
