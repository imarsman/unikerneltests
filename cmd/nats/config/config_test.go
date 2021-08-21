package config

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/matryer/is"
)

// TestCall test call for numbered cartoon
func TestCall(t *testing.T) {
	is := is.New(t)
	is.True(Config() != nil)
	c := Config()
	t.Log("config", spew.Sprintf("%+v", c))
}
