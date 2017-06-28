package govultr

import (
	"github.com/parnurzeal/gorequest"
	"encoding/json"
)

func AccountInfo(client *Client) (gorequest.Response, string, []error) {
	request := NewRequest(client.APIKey)
	return request.Get(API_URL + "v1/account/info").End()
}

func AuthInfo(client *Client) (gorequest.Response, string, []error) {
	request := NewRequest(client.APIKey)
	return request.Get(API_URL + "v1/auth/info").End()
}

func ListApps(client *Client) (gorequest.Response, string, []error) {
	request := NewRequest(client.APIKey)
	return request.Get(API_URL + "v1/app/list").End()
}

func ListBackups(client *Client) (gorequest.Response, string, []error) {
	request := NewRequest(client.APIKey)
	return request.Get(API_URL + "v1/backup/list").End()
}

func ListOS(client *Client) (map[string]OS, []error) {
	request := NewRequest(client.APIKey)
	_, body, errs := request.Get(API_URL + "v1/os/list").End()
	if len(errs) > 1 {
		return nil, errs
	} else {
		var oses map[string]OS
		err := json.Unmarshal([]byte(body), &oses)
		if err != nil {
			var errs []error
			return nil, append(errs, err)
		} else {
			return oses, nil
		}
	}
}