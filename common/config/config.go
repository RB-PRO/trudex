package config

import (
	"github.com/pkg/errors"
	"os"

	"gopkg.in/yaml.v2"
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
