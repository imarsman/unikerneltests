package instance

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/matryer/is"
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
