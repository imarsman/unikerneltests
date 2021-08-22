package instance

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/matryer/is"
)

// TestGCP test Google Cloud Platform instance
func TestGCP(t *testing.T) {
	is := is.New(t)
	client := NewForGCE()
	// is.True(client.Client != nil)
	t.Log("client", spew.Sprintf("%+v", client))
	ip, err := client.ExternalIP()
	t.Log("running in GCE", client.InCloud())
	is.True(err != nil)
	is.Equal(ip, "")
}
