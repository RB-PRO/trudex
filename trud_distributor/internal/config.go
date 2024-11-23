package internal

import (
	"context"
	"github.com/pkg/errors"
	"trudex/common/config"
)

type Config struct {
	Port int
}

func LoadConfigFromContext(ctx context.Context) (Config, error) {
	cfgVal := config.LoadFromContext(ctx)
	if cfgVal == nil {
		return Config{}, errors.New("error load config from context")
	}

	switch cfg := cfgVal.(type) {
	case Config:
		return cfg, nil
	default:
		return Config{}, errors.New("error parse config from context")
	}
}
