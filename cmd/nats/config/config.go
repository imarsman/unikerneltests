package config

import (
	// embed config file
	_ "embed"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var configBytes []byte

var config *Container

// Container config for app
type Container struct {
	Cloud     string   `yaml:"cloud" validate:"required,oneof=gcp"`
	ProjectID string   `yaml:"projectid" validate:"required"`
	Loging    *Logging `yaml:"logging" validate:"required,dive,required"`
}

// Logging settings tied to logging
type Logging struct {
	Name  string `yaml:"name" validate:"required"`
	Level string `yaml:"level" validate:"required,oneof=debug info alert warn error"`
}

// Config get loaded config
func Config() *Container {
	return config
}

func init() {
	config = &Container{}
	yaml.Unmarshal(configBytes, &config)

	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		panic(err)
	}
}
