package logging

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/logging"
	"github.com/imarsman/unikerneltests/cmd/nats/config"
)

// https://cloud.google.com/logging/docs/reference/libraries#linux-or-macos

var cloudLogger *logging.Logger
var logger *log.Logger
var client *logging.Client
var logName string

var debugLogger *log.Logger
var infoLogger *log.Logger
var alertLogger *log.Logger
var warnLogger *log.Logger
var errorLogger *log.Logger

func flush() {
	err := cloudLogger.Flush()
	if err != nil {
		fmt.Println("err", err)
	}
}

func init() {
	ctx := context.Background()

	// Use configured project ID for logs
	projectID := config.Config().ProjectID

	var err error
	// Creates a client.
	client, err = logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	logName = config.Config().LogName
	cloudLogger = client.Logger(logName)

	debugLogger = cloudLogger.StandardLogger(logging.Debug)
	infoLogger = cloudLogger.StandardLogger(logging.Info)
	alertLogger = cloudLogger.StandardLogger(logging.Alert)
	warnLogger = cloudLogger.StandardLogger(logging.Warning)
	errorLogger = cloudLogger.StandardLogger(logging.Error)

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		signal.Notify(stop, syscall.SIGTERM)

		// Wait for signal then close logging client
		<-stop

		client.Close()
	}()
}

// Debug make a debug log entry
func Debug(msg ...string) {
	debugLogger.Println(msg)
}

// Info make an info log entry
func Info(msg ...string) {
	infoLogger.Println(msg)
}

// Alert make an alert log entry
func Alert(msg ...string) {
	alertLogger.Println(msg)
}

// Warn log a warning entry
func Warn(msg ...string) {
	warnLogger.Println(msg)
}

// Error log an error entry
func Error(msg ...string) {
	errorLogger.Println(msg)
}
