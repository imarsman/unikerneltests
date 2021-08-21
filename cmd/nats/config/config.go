package config

import (
	// embed config file
	_ "embed"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var configBytes []byte

var config *Container

// Container config for app
type Container struct {
	ProjectID string `yaml:"projectid"`
	LogName   string `yaml:"logname"`
}

// Config get loaded config
func Config() *Container {
	return config
}

func init() {
	config = &Container{}
	yaml.Unmarshal(configBytes, &config)
}
