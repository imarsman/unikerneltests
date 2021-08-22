package instance

import (
	"errors"
	"net/http"

	"cloud.google.com/go/compute/metadata"
	"inet.af/netaddr"
)

// GCEClient represents a GCEClient instance
type GCEClient struct {
	Client *metadata.Client
	OnGCE  bool
}

// NewGCEClient get a new GPC instance
func NewGCEClient() IFInstance {
	gce := GCEClient{}
	gce.Client = metadata.NewClient(http.DefaultClient)
	gce.OnGCE = metadata.OnGCE()

	return &gce
}

func init() {
}

// InCloud is code running in cloud
func (c *GCEClient) InCloud() bool {
	return c.OnGCE
}

// ExternalIP use metadata API to get external IP
func (c *GCEClient) ExternalIP() (netaddr.IP, error) {
	if !c.OnGCE {
		return netaddr.IP{}, errors.New("Not running on GCE")
	}
	ipStr, err := c.Client.ExternalIP()
	if err != nil {
		return netaddr.IP{}, err
	}
	ip, err := netaddr.ParseIP(ipStr)
	if err != nil {
		return netaddr.IP{}, err
	}

	return ip, nil
}

// InternalIP use metadata API to get internal IP
func (c *GCEClient) InternalIP() (netaddr.IP, error) {
	if !c.OnGCE {
		return netaddr.IP{}, errors.New("Not running on GCE")
	}
	ipStr, err := c.Client.InternalIP()
	if err != nil {
		return netaddr.IP{}, err
	}
	ip, err := netaddr.ParseIP(ipStr)
	if err != nil {
		return netaddr.IP{}, err
	}

	return ip, nil
}

// InstanceID get instance ID
func (c *GCEClient) InstanceID() (string, error) {
	if !c.OnGCE {
		return "", errors.New("Not running on GCE")
	}

	return c.Client.InstanceID()
}

// InstanceName get instance name
func (c *GCEClient) InstanceName() (string, error) {
	if !c.OnGCE {
		return "", errors.New("Not running on GCE")
	}

	return c.Client.InstanceName()
}

// InGroup is instance in a managed instance group
func (c *GCEClient) InGroup() (bool, error) {
	if !c.OnGCE {
		return false, errors.New("Not running on GCE")
	}

	return false, nil
}

// GroupInstances get list of instances in MIG
func (c *GCEClient) GroupInstances() ([]*Instance, error) {
	if !c.OnGCE {
		return nil, errors.New("Not running on GCE")
	}

	return []*Instance{}, nil
}
