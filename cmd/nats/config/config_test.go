package config

import (
	"testing"

	"github.com/matryer/is"
)

// TestCall test call for numbered cartoon
func TestCall(t *testing.T) {
	is := is.New(t)
	is.True(Config() != nil)
	t.Logf("config %+v", Config())
}
