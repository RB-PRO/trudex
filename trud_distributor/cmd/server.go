package cmd

import (
	"context"
	"github.com/pkg/errors"
	"trudex/common/logger"
	"trudex/trud_distributor/internal/routers"
	"trudex/trud_distributor/internal/services"
	"trudex/trud_distributor/internal/services/rabbitmq"
)

func RunServer(ctx context.Context, opts ...ServerOption) (*Closer, error) {
	// closer object for exit from application
	closer := NewCloser()

	// option for start application
	cfgServer := &ServerOpts{}
	for _, opt := range opts {
		opt(cfgServer)
	}

	// load config
	lg := logger.NewLogger()

	rabbitmqService, stopFunc, err := rabbitmq.NewService()
	closer.Add(stopFunc)
	if err != nil {
		return closer, errors.Wrap(err, "failed to create rabbitmq service")
	}

	serviceList := services.NewServices(
		rabbitmqService,
	)

	// initial router
	engine := routers.InitRouter(lg, serviceList)

	// run router
	if err := engine.Run(":8083"); err != nil {
		return closer, errors.Wrap(err, "failed to start server")
	}

	return closer, nil
}

type ServerOpts struct {
	ConfigPatch string
}
type ServerOption func(*ServerOpts)

func WithConfigPatch(configPatch string) ServerOption {
	return func(h *ServerOpts) {
		h.ConfigPatch = configPatch
	}
}

func runRouter(ctx context.Context) {

}
