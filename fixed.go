package govultr

import (
	"github.com/parnurzeal/gorequest"
	"encoding/json"
)

func AccountInfo(client *Client) (gorequest.Response, string, []error) {
	return gorequest.New().Get(API_URL + "v1/account/info").Set("API-Key", client.APIKey).End()
}

func AuthInfo(client *Client) (gorequest.Response, string, []error) {
	return gorequest.New().Get(API_URL + "v1/auth/info").Set("API-Key", client.APIKey).End()
}

func ListApps(client *Client) (gorequest.Response, string, []error) {
	return gorequest.New().Get(API_URL + "v1/app/list").Set("API-Key", client.APIKey).End()
}

func ListBackups(client *Client) (gorequest.Response, string, []error) {
	return gorequest.New().Get(API_URL + "v1/backup/list").Set("API-Key", client.APIKey).End()
}

func ListOS(client *Client) (map[string]OS, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/os/list").Set("API-Key", client.APIKey).End()
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