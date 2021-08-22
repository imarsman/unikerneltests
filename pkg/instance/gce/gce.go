package gce

import (
	"errors"
	"net/http"

	"cloud.google.com/go/compute/metadata"
)

// Client represents a Client instance
type Client struct {
	Client *metadata.Client
	OnGCE  bool
}

// NewClient get a new GPC instance
func NewClient() *Client {
	gcp := Client{}
	gcp.Client = metadata.NewClient(http.DefaultClient)
	gcp.OnGCE = metadata.OnGCE()

	return &gcp
}

func init() {
}

// ExternalIP use metadata API to get external IP
func (c *Client) ExternalIP() (string, error) {
	if !c.OnGCE {
		return "", errors.New("Not running on GCE")
	}
	ip, err := c.Client.ExternalIP()
	if err != nil {
		return "", err
	}

	return ip, nil
}

// InternalIP use metadata API to get internal IP
func (c *Client) InternalIP() (string, error) {
	if !c.OnGCE {
		return "", errors.New("Not running on GCE")
	}
	ip, err := c.Client.InternalIP()
	if err != nil {
		return "", err
	}

	return ip, nil
}

// InGroup is instance in a managed instance group
func (c *Client) InGroup() (bool, error) {
	if !c.OnGCE {
		return false, errors.New("Not running on GCE")
	}

	return false, nil
}

// GroupInstances get list of instances in MIG
func (c *Client) GroupInstances() ([]string, error) {
	if !c.OnGCE {
		return nil, errors.New("Not running on GCE")
	}

	return []string{"hello"}, nil
}
