package main

import (
	"fmt"

	"github.com/imarsman/unikerneltests/cmd/nats/cloudlog"
	"github.com/imarsman/unikerneltests/cmd/nats/config"
	"github.com/imarsman/unikerneltests/pkg/instance"
	"github.com/nats-io/nats-server/v2/server"
	stand "github.com/nats-io/nats-streaming-server/server"
	"github.com/nats-io/nats.go"
)

// Basic NATS server setup. Plans are to allow for standalone and clustered NATS
// server setups.

func main() {
	snopts := stand.NewNATSOptions()
	snopts.Port = nats.DefaultPort
	snopts.HTTPPort = 8223

	// Now run the server with the streaming and streaming/nats options.
	ns, err := server.NewServer(snopts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Starting NAT server on %v\n", nats.DefaultPort)
	// Start things up. Block here until done.
	if err := server.Run(ns); err != nil {
		server.PrintAndDie(err.Error())
	}

	if config.Config().Cloud.Type == config.CloudGCE {
		client := instance.NewGCEClient()
		externalIP, err := client.ExternalIP()
		if err != nil {
			cloudlog.Info("Instance IP", externalIP.String(), "error", err.Error())
		}
	}

	ns.WaitForShutdown()
}
