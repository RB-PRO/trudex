package rabbitmq

import (
	"context"
)

type ServiceInterface interface {
	Push(ctx context.Context, params ...interface{}) (bool, error)
}
