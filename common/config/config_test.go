package config_test

import (
	"trudex/common/config"

	"testing"
)

func TestConfig(t *testing.T) {
	const configFilePatch = "./test_config.yaml"

	type TestConfig struct {
		Port        string   `json:"port"`
		Stars       []string `json:"stars"`
		StructParam struct {
			OneParam string `json:"one_param"`
			TwoParam string `json:"two_param"`
		} `json:"struct_param"`
	}
	testData := `
port: 8080
stars: [one,two,three]
struct_param:
  one_param: "one"
  two_param: "two"
`

	configService, err := config.New[TestConfig](
		config.WithData(testData),
	)
	if err != nil {
		t.Fatal(err)
	}
	if configService == nil {
		t.Log(configService)
		t.Fatal("configService in nil")
	}

	z := configService.Config()
	_ = z
}
