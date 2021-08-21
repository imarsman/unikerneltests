package instance

import (
	"net/http"

	"github.com/imarsman/unikerneltests/pkg/instance/metadata"
)

var client *metadata.Client

func init() {
	client = metadata.NewClient(http.DefaultClient)
}

// ExternalIP use metadata API to get external IP
func ExternalIP() (string, error) {
	ip, err := client.ExternalIP()
	if err != nil {
		return "", err
	}
	return ip, nil
}
