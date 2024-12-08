package config_test

import (
	"trudex/common/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	type TestConfig struct {
		Port        string   `yaml:"port"`
		Stars       []string `yaml:"stars"`
		StructParam struct {
			OneParam string `yaml:"one_param"`
			TwoParam string `yaml:"two_param"`
		} `yaml:"struct_param"`
	}

	testData := `
port: 8080
stars: [one,two,three]
struct_param:
  one_param: "one"
  two_param: "two"
`

	expectedData := TestConfig{
		Port:  "8080",
		Stars: []string{"one", "two", "three"},
		StructParam: struct {
			OneParam string `yaml:"one_param"`
			TwoParam string `yaml:"two_param"`
		}{OneParam: "one", TwoParam: "two"},
	}

	configService, err := config.New[TestConfig](
		config.WithData(testData),
	)
	assert.NoError(t, err)
	assert.Nil(t, configService)

	cfg := configService.Config()
	assert.Equal(t, expectedData, cfg)
}
