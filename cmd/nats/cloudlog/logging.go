package cloudlog

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"cloud.google.com/go/logging"
	"github.com/imarsman/unikerneltests/cmd/nats/config"
)

// https://cloud.google.com/logging/docs/reference/libraries#linux-or-macos

const (
	levelDebug = iota
	levelInfo
	levelAlert
	levelWarn
	levelError
)

const (
	debugLevelName = "debug"
	infoLevelName  = "info"
	alertLevelName = "alert"
	warnLevelName  = "warn"
	errorLevelName = "error"
)

var cloudLogger *logging.Logger
var logger *log.Logger
var client *logging.Client
var logName string
var logLevel int

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
	projectID := config.Config().Cloud.ProjectID

	var err error
	// Creates a client.
	client, err = logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	logName = config.Config().Loging.Name
	cloudLogger = client.Logger(logName)
	level := config.Config().Loging.Level
	logLevel = 0

	switch strings.ToLower(level) {
	case debugLevelName:
		logLevel = levelDebug
	case infoLevelName:
		logLevel = levelInfo
	case alertLevelName:
		logLevel = levelAlert
	case warnLevelName:
		logLevel = levelWarn
	case errorLevelName:
		logLevel = levelError
	}

	// Set up loggers for different log levels
	debugLogger = cloudLogger.StandardLogger(logging.Debug)
	infoLogger = cloudLogger.StandardLogger(logging.Info)
	alertLogger = cloudLogger.StandardLogger(logging.Alert)
	warnLogger = cloudLogger.StandardLogger(logging.Warning)
	errorLogger = cloudLogger.StandardLogger(logging.Error)

	// Wait for signal and exit cleanly
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
	if logLevel >= levelDebug { // Check log level otherwise do nothing
		debugLogger.Println(msg)
	}
}

// Info make an info log entry
func Info(msg ...string) {
	if logLevel >= levelInfo { // Check log level otherwise do nothing
		infoLogger.Println(msg)
	}
}

// Alert make an alert log entry
func Alert(msg ...string) {
	if logLevel >= levelAlert { // Check log level otherwise do nothing
		alertLogger.Println(msg)
	}
}

// Warn log a warning entry
func Warn(msg ...string) {
	if logLevel >= levelWarn { // Check log level otherwise do nothing
		warnLogger.Println(msg)
	}
}

// Error log an error entry
func Error(msg ...string) {
	if logLevel >= levelError { // Check log level otherwise do nothing
		errorLogger.Println(msg)
	}
}
