package instance

import (
	"errors"
	"fmt"
	"net/http"

	"cloud.google.com/go/compute/metadata"
	"inet.af/netaddr"
)

// GCE specific function implementation

// GCEClient represents a GCEClient instance
type GCEClient struct {
	Client *metadata.Client
	OnGCE  bool
}

// NewGCEClient get a new GCE instance as an IFInstance
func NewGCEClient() IFInstance {
	gce := GCEClient{}

	gce.Client = metadata.NewClient(http.DefaultClient)
	if gce.Client == nil {
		fmt.Println("metadata client", gce.Client)
	}
	gce.OnGCE = metadata.OnGCE()

	return &gce
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
		fmt.Println("Error getting external IP")
		return netaddr.IP{}, err
	}

	fmt.Println("ip string", ipStr)

	ip, err := netaddr.ParseIP(ipStr)
	if err != nil {
		return netaddr.IP{}, err
	}

	fmt.Println("ip", ip.String())

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
func (c *GCEClient) InGroup(name string) (bool, error) {
	if !c.OnGCE {
		return false, errors.New("Not running on GCE")
	}

	return false, nil
}

// GroupInstances get list of instances in MIG
func (c *GCEClient) GroupInstances(name string) ([]*Instance, error) {
	group := NewGroupInstances()

	if !c.OnGCE {
		return group, errors.New("Not running on GCE")
	}

	return group, nil
}
