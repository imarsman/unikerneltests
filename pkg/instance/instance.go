package instance

import (
	"fmt"

	// Lilely to be added to Go
	// https://github.com/golang/go/issues/46518
	"inet.af/netaddr"
)

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
