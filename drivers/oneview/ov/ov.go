package ov

import (
	"github.com/docker/machine/drivers/oneview/rest"
)
// OVClient - wrapper class for ov api's
type OVClient struct {
	rest.Client
}

// new Client
func (c *OVClient) NewOVClient(user string, password string, domain string, endpoint string, sslverify bool, apiversion int) (*OVClient) {
	return &OVClient{
		rest.Client{
			User:       user,
			Password:   password,
			Domain:     domain,
			Endpoint:   endpoint,
			SSLVerify:  sslverify,
			APIVersion: apiversion,
			APIKey:     "none",
		},
	}
}
