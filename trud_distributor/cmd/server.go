package cmd

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"trudex/common/config"
	"trudex/common/logger"
	appconfig "trudex/trud_distributor/internal"
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
	ctx, err := config.AddConfigToCtx[appconfig.Config](ctx, cfgServer.ConfigPatch)
	if err != nil {
		return closer, errors.Wrap(err, "failed to load and add to context config")
	}

	// init logger
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
	srv := routers.InitRouter(ctx, lg, serviceList)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	closer.Add(func(ctx context.Context) error {
		// Shutdown the server gracefully
		if err := srv.Shutdown(ctx); err != nil {
			return errors.Wrap(err, "server shutdown failed")
		}

		return nil
	})

	// Create a channel to listen for interrupt signals
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

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
