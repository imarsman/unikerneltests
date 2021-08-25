package instance

import (
	// Lilely to be added to Go
	// https://github.com/golang/go/issues/46518
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/compute/metadata"
	"inet.af/netaddr"
)

// A basic set of functions that can be implemented for various cloud platforms
// Each cloud provider makes available many pieces of information about
// instances and their contexts. This is not the goal here. We mostly need to
// know things like internal and external IP, and instance summary information
// for instance groups.

// WaitForMetadataService test to see if API can be used
func WaitForMetadataService() error {
	c := make(chan int)

	var err error
	go func() {
		for i := 0; i < 10; i++ {
			t := time.Now()
			client := metadata.NewClient(http.DefaultClient)
			d := time.Since(t)
			_, err := client.ProjectID()
			if err != nil {
				fmt.Println("Waiting")
				time.Sleep(2 * time.Second)
			} else {
				fmt.Println("Finished wait. Took", d.Round(time.Nanosecond), "to get metadata client")
				c <- 1
				return
			}
		}
		c <- 1
	}()

	<-c
	return err
}

// IFInstance a cross-cloud instance interface
type IFInstance interface {
	// External IP for an instance, if running in cloud
	ExternalIP() (netaddr.IP, error)
	// Internal IP for an instance, if running in cloud
	InternalIP() (netaddr.IP, error)
	// Is container running in a group
	InGroup(string) (bool, error)
	// Instances in group, if any
	GroupInstances(string) ([]*Instance, error)
	// Id of running instance
	InstanceID() (string, error)
	// Name of running instance
	InstanceName() (string, error)
	// Is code running in a cloud
	InCloud() bool
}

// Instance useful information about an instance across platform
type Instance struct {
	PublicIP     netaddr.IP
	PrivateIP    netaddr.IP
	Zone         string
	ProjectID    string
	CreationDate string
}

// NewGroupInstances get empty list of group instances
func NewGroupInstances() []*Instance {
	group := []*Instance{}
	i := new(Instance)
	group = append(group, i)

	return group
}

func init() {
	// ip, _ := netaddr.ParseIP("127.0.0.1")
	// fmt.Println("ip", ip.String())
}
