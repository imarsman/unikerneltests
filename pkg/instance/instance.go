package instance

import (
	"fmt"

	// Lilely to be added to Go
	// https://github.com/golang/go/issues/46518
	"inet.af/netaddr"
)

// A basic set of functions that can be implemented for various cloud platforms
// Each cloud provider makes available many pieces of information about
// instances and their contexts. This is not the goal here. We mostly need to
// know things like internal and external IP, and instance summary information
// for instance groups.

// IFInstance a cross-cloud instance interface
type IFInstance interface {
	ExternalIP() (netaddr.IP, error)
	InternalIP() (netaddr.IP, error)
	InGroup() (bool, error)
	GroupInstances() ([]*Instance, error)
	InstanceID() (string, error)
	InstanceName() (string, error)
	InCloud() bool
}

// Instance useful information about an instance across platform
type Instance struct {
	PublicIPs    []netaddr.IP
	PrivateIPs   []netaddr.IP
	Zone         string
	ProjectID    string
	CreationDate string
}

func init() {
	ip, _ := netaddr.ParseIP("127.0.0.1")
	fmt.Println("ip", ip.String())
}
