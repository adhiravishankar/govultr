package govultr

import (
	"github.com/parnurzeal/gorequest"
)

func AccountInfo(client *Client) (Account, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/account/info").Set("API-Key", client.APIKey).End()
	var account Account
	errs2 := UnmarshalJson(body, errs, &account)
	return account, errs2
}

func AuthInfo(client *Client) (Auth, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/auth/info").Set("API-Key", client.APIKey).End()
	var auth Auth
	errs2 := UnmarshalJson(body, errs, &auth)
	return auth, errs2
}

func ListApps(client *Client) (map[string]App, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/app/list").Set("API-Key", client.APIKey).End()
	var apps map[string]App
	errs2 := UnmarshalJson(body, errs, &apps)
	return apps, errs2
}

func ListBackups(client *Client) (map[string]Backup, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/backup/list").Set("API-Key", client.APIKey).End()
	var backups map[string]Backup
	errs2 := UnmarshalJson(body, errs, &backups)
	return backups, errs2
}

func ListOS(client *Client) (map[string]OS, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "v1/os/list").Set("API-Key", client.APIKey).End()
	var oses map[string]OS
	errs2 := UnmarshalJson(body, errs, &oses)
	return oses, errs2
}