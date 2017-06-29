package govultr

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

var API_URL = "https://api.vultr.com/"

type Client struct {
	APIKey string
}

//noinspection GoUnusedExportedFunction
func NewClient(key string) (*Client, error) {
	if len(key) == 0 {
		return nil, errors.New("vultr: client")
	}
	return &Client{ APIKey: key }, nil
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

func NoResponse(response gorequest.Response, errs []error) (bool, []error) {
	if len(errs) > 0 {
		return false, errs
	} else if response.StatusCode == 200 {
		return true, nil
	} else {
		return false, nil
	}
}