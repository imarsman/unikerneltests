package main

import (
	"fmt"

	"github.com/imarsman/unikerneltests/cmd/nats/cloudlog"
	"github.com/nats-io/nats-server/v2/server"
	stand "github.com/nats-io/nats-streaming-server/server"
	"github.com/nats-io/nats.go"
)

// Basic NATS server setup. Plans are to allow for standalone and clustered NATS
// server setups.

// Leader choosing library
// https://github.com/nats-io/graft

func main() {
	// Need to
	// - check to see if running in group
	// - if running in group, get IP
	// - Choose leader from among servers
	// - Use leader IP to set up NATS in cluster mode
	// - Run NATS

	natsOpts := stand.NewNATSOptions()
	natsOpts.Port = nats.DefaultPort
	// snopts.HTTPPort = 8222

	// Now run the server with the streaming and streaming/nats options.
	natsServer, err := server.NewServer(natsOpts)
	if err != nil {
		panic(err)
	}

	// Start things up. Block here until done.
	if err := server.Run(natsServer); err != nil {
		server.PrintAndDie(err.Error())
	}

	// if config.Config().Cloud.Type == config.CloudGCE {
	// 	client := instance.NewGCEClient()

	// 	if client.InCloud() {
	// 		cloudlog.Info("Starting NAT server on", nats.DefaultPort)
	// 	} else {
	// 		fmt.Println("Starting NAT server on", nats.DefaultPort)
	// 	}

	// 	externalIP, err := client.ExternalIP()
	// 	if err == nil {
	// 		cloudlog.Info("Instance IP", externalIP.String(), "error", err.Error())
	// 	} else {
	// 		fmt.Println("No IP found", err.Error())
	// 	}
	// }
	cloudlog.Info("Hello from VM")
	fmt.Println("Logging to syslog test")

	natsServer.WaitForShutdown()
}
