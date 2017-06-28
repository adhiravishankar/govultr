package govultr

import (
	"github.com/parnurzeal/gorequest"
)

var API_URL = "https://api.vultr.com/"

type Client struct {
	APIKey string
}

func newClient(key string) *Client {
	return &Client{ APIKey: key }
}

func newRequest(key string) *gorequest.SuperAgent {
	request := gorequest.New()
	request.Header = map[string]string{"API-Key": key}
	return request
}