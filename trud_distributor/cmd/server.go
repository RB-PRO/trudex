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
	"trudex/trud_distributor/internal"
	"trudex/trud_distributor/internal/routers"
	"trudex/trud_distributor/internal/services"
	"trudex/trud_distributor/internal/services/rabbitmq"
)

//func readConfigToCtx(ctx context.Context, patch string) (context.Context, error) {
//	cfgConfig, err := config.LoadConfig(patch)
//	if err != nil {
//		return ctx, err
//	}
//
//	ctx, err = config.CfgToContext(ctx, cfgConfig)
//
//	return config.CfgToContext(ctx, patch), nil
//}

const (
	defaultConfigPatch = "trud_distributor/trud_distributor.yaml"
)

func RunServer(ctx context.Context) (*Closer, error) {
	// closer object for exit from application
	closer := NewCloser()

	// initial logger
	lg := logger.NewLogger()

	// initial config service
	configPatch := defaultConfigPatch
	if customConfigPatch := os.Getenv("CUSTOM_CONFIG_PATCH"); customConfigPatch != "" {
		configPatch = customConfigPatch
	}

	configService, err := config.New[internal.Config](
		config.WithConfigPatch(configPatch),
	)
	if err != nil {
		return closer, errors.Wrap(err, "failed to create config service")
	}

	// initial rabbitmq service
	rabbitmqService, stopFunc, err := rabbitmq.NewService()
	closer.Add(stopFunc)
	if err != nil {
		return closer, errors.Wrap(err, "failed to create rabbitmq service")
	}

	serviceList := services.NewServices(
		configService,
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
