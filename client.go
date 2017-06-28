package govultr

import "encoding/json"

var API_URL = "https://api.vultr.com/"

type Client struct {
	APIKey string
}

func NewClient(key string) *Client {
	return &Client{ APIKey: key }
}

func UnmarshalJson(body string, errs []error, entity interface{}) []error  {
	if len(errs) > 1 {
		return errs
	} else {
		err := json.Unmarshal([]byte(body), &entity)
		if err != nil {
			var errs []error
			return append(errs, err)
		} else {
			return nil
		}
	}
}