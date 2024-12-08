package config_test

import (
	"context"
	"trudex/common/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	type Another struct {
		Val string `yaml:"val"`
	}

	type TestConfig struct {
		Port        string   `yaml:"port"`
		Stars       []string `yaml:"stars"`
		StructParam struct {
			OneParam string `yaml:"one_param"`
			TwoParam string `yaml:"two_param"`
		} `yaml:"struct_param"`
		Another Another `yaml:"another"`
	}

	testData := `
port: 8080
stars: [one,two,three]
struct_param:
  one_param: "one"
  two_param: "two"
another:
  val: "val"
`

	expectedData := TestConfig{
		Port:  "8080",
		Stars: []string{"one", "two", "three"},
		StructParam: struct {
			OneParam string `yaml:"one_param"`
			TwoParam string `yaml:"two_param"`
		}{OneParam: "one", TwoParam: "two"},
		Another: Another{Val: "val"},
	}

	configService, err := config.New[TestConfig](
		config.WithData(testData),
	)
	assert.NoError(t, err)
	assert.NotNil(t, configService)

	cfg := configService.Config()
	assert.Equal(t, expectedData, cfg)

	// load substructure
	ctx := context.Background()
	keys := make(map[string]any)
	keys[config.CtxKey] = cfg
	ctx, _, err = config.LoadToCtxFromKeys[TestConfig](ctx, keys)
	assert.NoError(t, err)

	loadCfg := config.LoadFromCtx[TestConfig](ctx)
	assert.Equal(t, expectedData, loadCfg)

	loadAnotherCfg := config.LoadFromCtx[Another](ctx)
	assert.Equal(t, Another{Val: "val"}, loadAnotherCfg)

}
