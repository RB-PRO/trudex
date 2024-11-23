package config

import (
	"context"
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type cfgRegistry = struct{}

type config struct {
	cfg *koanf.Koanf
}

func AddConfigToCtx[T any](ctx context.Context, patch string) (context.Context, error) {
	cfg := koanf.New(".")

	if err := cfg.Load(file.Provider(patch), yaml.Parser()); err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, cfgRegistry{}, config{cfg}), nil
}

func LoadFromContext(ctx context.Context) any {
	z := ctx.Value(cfgRegistry{})
	if z != nil {
		x := z.(config)

		_ = x.cfg
		y := x.cfg.Get("rabbitmq")
		_ = y
		fmt.Println("$$#")
	}

	return ctx.Value(cfgRegistry{}).(config).cfg
}
