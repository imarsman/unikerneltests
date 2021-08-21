package instance

import (
	"net/http"

	"cloud.google.com/go/compute/metadata"
)

var client *metadata.Client

// This package provides GCP specific metadata

func init() {
	client = metadata.NewClient(http.DefaultClient)
}

// ExternalIP use metadata API to get external IP
func ExternalIP() (string, error) {
	ip, err := client.ExternalIP()
	if err != nil {
		return "", err
	}
	return ip, nil
}

// InternalIP use metadata API to get internal IP
func InternalIP() (string, error) {
	ip, err := client.InternalIP()
	if err != nil {
		return "", err
	}
	return ip, nil
}
