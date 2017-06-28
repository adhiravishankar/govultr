package govultr

import (
	"github.com/parnurzeal/gorequest"
)

func AccountInfo(client *Client) (gorequest.Response, string, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/account/info").Set("API-Key", client.APIKey).End()
}

func AuthInfo(client *Client) (Auth, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/auth/info").Set("API-Key", client.APIKey).End()
	var auth Auth
	errs2 := UnmarshalJson(body, errs, &auth)
	return auth, errs2
}

func ListApps(client *Client) (gorequest.Response, string, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/app/list").Set("API-Key", client.APIKey).End()
}

func ListBackups(client *Client) (gorequest.Response, string, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/backup/list").Set("API-Key", client.APIKey).End()
}

func ListOS(client *Client) (map[string]OS, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/os/list").Set("API-Key", client.APIKey).End()
	var oses map[string]OS
	errs2 := UnmarshalJson(body, errs, &oses)
	return oses, errs2
}