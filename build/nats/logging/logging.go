package logging

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/logging"
	"github.com/imarsman/unikerneltests/cmd/nats/config"
)

// https://cloud.google.com/logging/docs/reference/libraries#linux-or-macos

var logger *log.Logger
var client *logging.Client
var logName string

// Logger get reference to logger
func Logger() *log.Logger {
	return logger
}

func init() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := config.Config().ProjectID

	var err error
	// Creates a client.
	client, err = logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name of the log to write to.
	logName = config.Config().LogName

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	<-stop

	client.Close()
}

// Debug make a debug log entry
func Debug(msg ...string) {
	logger = client.Logger(logName).StandardLogger(logging.Debug)
	logger.Println(msg)
}

// Info make an info log entry
func Info(msg ...string) {
	logger = client.Logger(logName).StandardLogger(logging.Info)
	logger.Println(msg)
}

// Alert make an alert log entry
func Alert(msg ...string) {
	logger = client.Logger(logName).StandardLogger(logging.Alert)
	logger.Println(msg)
}

// Warn log a warning entry
func Warn(msg ...string) {
	logger = client.Logger(logName).StandardLogger(logging.Warning)
	logger.Println(msg)
}

// Error log an error entry
func Error(msg ...string) {
	logger = client.Logger(logName).StandardLogger(logging.Error)
	logger.Println(msg)
}
