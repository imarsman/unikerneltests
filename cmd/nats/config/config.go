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
	Cloud   *Cloud   `yaml:"cloud" validate:"required"`
	Loging  *Logging `yaml:"logging" validate:"required,dive,required"`
	Context *Context `yaml:"context" validate:"required"`
}

// Cloud configuration for cloud used
type Cloud struct {
	Type      string `yaml:"type" validate:"required,oneof=gce"`
	ProjectID string `yaml:"projectid" validate:"required"`
}

// Logging settings tied to logging
type Logging struct {
	Name  string `yaml:"name" validate:"required"`
	Level string `yaml:"level" validate:"required,oneof=debug info alert warn error"`
}

// Context the context an instance is running in
type Context struct {
	Instance  string `yaml:"instance" validate:"required,oneof=standalone group"`
	GroupName string `yaml:"groupname"`
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
