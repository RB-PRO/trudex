// Package config Package for holding the config and accessing it from outside
//
// # Very convenient to request sub-configs from different parts of the program and work with context
//
// Restriction: If you want to get a sub-config, then its name and the type of its structure must match
package config

import (
	"context"
	"github.com/pkg/errors"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

type Service[T any] struct {
	configType T
}

func New[T any](opts ...OptsFunc) (*Service[T], error) {
	cfgForConfig := &Option{}
	for _, opt := range opts {
		opt(cfgForConfig)
	}

	var cfg T
	var data []byte

	if cfgForConfig.Data != "" {
		data = []byte(cfgForConfig.Data)
	}

	if cfgForConfig.ConfigPatch != "" {
		var err error
		data, err = os.ReadFile(cfgForConfig.ConfigPatch)
		if err != nil {
			return nil, errors.Wrapf(err, "could not read error file: '%s'", cfgForConfig.ConfigPatch)
		}
	}

	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "could not unmarshal the config: '%s'", string(data))
	}

	return &Service[T]{cfg}, nil
}

const (
	CtxKey = "cfg"
)

func LoadFromCtx[T any](ctx context.Context) T {
	val := ctx.Value(CtxKey)
	if val == nil {
		return *new(T) // omg
	}

	z, ok := val.(T)
	if !ok {
		// if the requested value is not the root value in the structure,
		// then we look inside the structure under investigation
		typeName := reflect.TypeOf(*new(T)).Name()

		// todo: the name of the structure and its definition in the parent structure must match. Fix this shitty code.

		fields := reflect.ValueOf(val).FieldByName(typeName)
		if !fields.IsValid() {
			return *new(T)
		}

		return fields.Interface().(T)
	}
	return z
}

func LoadToCtx[T any](ctx context.Context, cfg T) context.Context {
	return context.WithValue(ctx, CtxKey, cfg)
}

func LoadToCtxFromKeys[T any](ctx context.Context, keys map[string]any) (context.Context, T, error) {
	conf := keys[CtxKey]
	if conf == nil {
		return ctx, *new(T), errors.Errorf("no config found for the key: %s", CtxKey)
	}

	switch cfg := conf.(type) {
	case T:
		return context.WithValue(ctx, CtxKey, cfg), cfg, nil
	default:
		return ctx, *new(T), errors.Errorf("invalid config type for the key: %s", CtxKey)
	}
}

func (s *Service[T]) Config() T {
	return s.configType
}

// Option generate optionals params for server
type Option struct {
	Data        string
	ConfigPatch string
}

type OptsFunc func(*Option)

func WithConfigPatch(configPatch string) OptsFunc {
	return func(h *Option) {
		h.ConfigPatch = configPatch
	}
}

func WithData(configPatch string) OptsFunc {
	return func(h *Option) {
		h.Data = configPatch
	}
}
