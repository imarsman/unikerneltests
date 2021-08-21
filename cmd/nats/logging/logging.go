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

// Logger get reference to logger
// func Logger() *log.Logger {
// 	return logger
// }

func flush() {
	err := cloudLogger.Flush()
	if err != nil {
		fmt.Println("err", err)
	}
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

	// ctx2 := context.Background()
	// err = client.Ping(ctx2)
	// fmt.Println("error", err)
	// Sets the name of the log to write to.
	logName = config.Config().LogName
	cloudLogger = client.Logger(logName)

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		signal.Notify(stop, syscall.SIGTERM)

		fmt.Println("waiting")
		<-stop

		client.Close()
	}()
}

// Debug make a debug log entry
func Debug(msg ...string) {
	stdL := cloudLogger.StandardLogger(logging.Debug)
	stdL.Println(msg)
	// cloudLogger.Flush()
}

// Info make an info log entry
func Info(msg ...string) {
	stdL := cloudLogger.StandardLogger(logging.Info)
	stdL.Println(msg)
	// cloudLogger.Flush()
}

// Alert make an alert log entry
func Alert(msg ...string) {
	stdL := cloudLogger.StandardLogger(logging.Alert)
	stdL.Println(msg)
	// cloudLogger.Flush()
}

// Warn log a warning entry
func Warn(msg ...string) {
	stdL := cloudLogger.StandardLogger(logging.Warning)
	stdL.Println(msg)
	// cloudLogger.Flush()
}

// Error log an error entry
func Error(msg ...string) {
	stdL := cloudLogger.StandardLogger(logging.Error)
	stdL.Println(msg)
	// cloudLogger.Flush()
}
