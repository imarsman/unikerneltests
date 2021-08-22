package instance

import (
	"errors"
	"net/http"

	"cloud.google.com/go/compute/metadata"
	"github.com/imarsman/unikerneltests/pkg/instance/gce"
)

// Instance a cross-cloud instance interface
type Instance interface {
	ExternalIP() (string, error)
	InternalIP() (string, error)
	InGroup() (bool, error)
	GroupInstances() ([]string, error)
	InCloud() bool
}

// NewForGCE get new instance for GCE environment
func NewForGCE() Instance {
	client := gce.NewGCEClient()

	return client
}

// GCEClient represents a GCEClient instance
type GCEClient struct {
	Client *metadata.Client
	OnGCE  bool
}

// NewGCEClient get a new GPC instance
func NewGCEClient() *GCEClient {
	gcp := GCEClient{}
	gcp.Client = metadata.NewClient(http.DefaultClient)
	gcp.OnGCE = metadata.OnGCE()

	return &gcp
}

func init() {
}

// InCloud is code running in cloud
func (c *GCEClient) InCloud() bool {
	return c.OnGCE
}

// ExternalIP use metadata API to get external IP
func (c *GCEClient) ExternalIP() (string, error) {
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
func (c *GCEClient) InternalIP() (string, error) {
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
func (c *GCEClient) InGroup() (bool, error) {
	if !c.OnGCE {
		return false, errors.New("Not running on GCE")
	}

	return false, nil
}

// GroupInstances get list of instances in MIG
func (c *GCEClient) GroupInstances() ([]string, error) {
	if !c.OnGCE {
		return nil, errors.New("Not running on GCE")
	}

	return []string{"hello"}, nil
}
