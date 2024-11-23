package rabbitmq

import (
	"context"
	"trudex/common/config"
)

type ServiceInterface interface {
	Push(ctx context.Context, params ...interface{}) (bool, error)
}

type Service struct{}

func NewService() (*Service, func(ctx context.Context) error, error) {
	stopFunc := func(ctx context.Context) error {
		// todo: implement me
		return nil
	}
	return &Service{}, stopFunc, nil
}

func (s *Service) Push(ctx context.Context, params ...interface{}) (bool, error) {

	zzz := config.LoadFromContext(ctx)
	_ = zzz

	return false, nil
}
