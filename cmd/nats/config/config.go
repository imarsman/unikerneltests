package config

import (
	// embed config file
	_ "embed"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

// Configuration for the app. It is useful to embed this at build time.

const (
	// CloudGCE GCE cloud
	CloudGCE = "gce"
	// CloudAWS AWS cloud
	CloudAWS = "aws"
)

//go:embed config.yaml
var configBytes []byte

var config *Container

// Container config for app
type Container struct {
	// Cloud specific settings
	Cloud *Cloud `yaml:"cloud" validate:"required"`
	// Logging settings
	Loging *Logging `yaml:"logging" validate:"required,dive,required"`
	// Context of running instance
	Context *Context `yaml:"context" validate:"required"`
}

// Cloud configuration for cloud used
type Cloud struct {
	// Cloud type - gce at the moment
	Type string `yaml:"type" validate:"required,oneof=gce"`
	// Project ID - required for now though not common to AWS
	ProjectID string `yaml:"projectid" validate:"required"`
}

// Logging settings tied to logging
type Logging struct {
	// Name of log for cloud logging
	Name string `yaml:"name" validate:"required"`
	// Default log level
	Level string `yaml:"level" validate:"required,oneof=debug info alert warn error"`
}

// Context the context an instance is running in
type Context struct {
	// Context instance is running in - standalone or group
	RunContext string `yaml:"runcontext" validate:"required,oneof=standalone group"`
	// The name of the instance group, if any, that instance is running in
	GroupName string `yaml:"groupname"`
}

// Config get loaded config
func Config() *Container {
	return config
}

// Initialize configuration and bail if there is a problem. Should show error
// prior to deployment as the errors are more likely to be tied to validation.
func init() {
	config = &Container{}
	yaml.Unmarshal(configBytes, &config)

	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		panic(err)
	}
}
