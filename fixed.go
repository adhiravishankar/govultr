package govultr

import "github.com/parnurzeal/gorequest"

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

func ListOS(client *Client) (gorequest.Response, string, []error) {
	request := NewRequest(client.APIKey)
	return request.Get(API_URL + "v1/os/list").End()
}