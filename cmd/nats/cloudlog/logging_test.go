package cloudlog

import (
	"testing"
)

// TestCall test call for numbered cartoon
func TestCall(t *testing.T) {
	// is := is.New(t)
	Debug("hello debug log")
	Info("hello info log")
	Warn("hello warn log")
	Error("hello error log")
	flush()
}
