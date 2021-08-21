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
	Cloud     string   `yaml:"cloud"`
	ProjectID string   `yaml:"projectid"`
	Loging    *Logging `yaml:"logging"`
}

// Logging settings tied to logging
type Logging struct {
	Name  string `yaml:"name"`
	Level string `yaml:"level"`
}

// Config get loaded config
func Config() *Container {
	return config
}

func init() {
	config = &Container{}
	yaml.Unmarshal(configBytes, &config)
}
