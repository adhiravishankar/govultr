package govultr

var API_URL = "https://api.vultr.com/"

type Client struct {
	APIKey string
}

func NewClient(key string) *Client {
	return &Client{ APIKey: key }
}