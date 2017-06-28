package govultr

import (
	"github.com/parnurzeal/gorequest"
)

var API_URL = "https://api.vultr.com/"

type Client struct {
	APIKey string
}

func NewClient(key string) *Client {
	return &Client{ APIKey: key }
}

func NewRequest(key string) *gorequest.SuperAgent {
	return gorequest.New().Set("API-Key", key)
}