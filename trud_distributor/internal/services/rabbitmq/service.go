package rabbitmq

import (
	"context"
)

type Service struct{}

func NewService() (*Service, func(ctx context.Context) error, error) {
	stopFunc := func(ctx context.Context) error {
		// todo: implement me
		return nil
	}
	return &Service{}, stopFunc, nil
}
