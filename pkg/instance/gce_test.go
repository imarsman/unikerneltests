package instance

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/matryer/is"
	"inet.af/netaddr"
)

// TestGCP test Google Cloud Platform instance
func TestGCP(t *testing.T) {
	is := is.New(t)
	client := NewGCEClient()

	t.Log("client", spew.Sprintf("%+v", client))
	externalIP, err := client.ExternalIP()
	t.Log("running in GCE", client.InCloud())

	externalIP, err = client.ExternalIP()
	is.True(err != nil)
	t.Log("External IP", externalIP, "err", err.Error())
	is.True(err != nil)

	is.True(externalIP.IsValid() == false)
}

func TestIP(t *testing.T) {
	is := is.New(t)

	ip, err := netaddr.ParseIP("35.196.70.51")
	is.NoErr(err)
	t.Log("ip", ip.String())
	// if err != nil {
	// 	return netaddr.IP{}, err
	// }
}
