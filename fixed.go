package govultr

import (
	"fmt"
)

func AccountInfo(client *Client) {
	request := newRequest(client.APIKey)
	resp, body , err := request.Get(API_URL + "v1/account/info").End()
	fmt.Println(resp, body, err)
}

func AuthInfo(client *Client) {
	request := newRequest(client.APIKey)
	resp, body , err := request.Get(API_URL + "v1/auth/info").End()
	fmt.Println(resp, body, err)
}

func ListApps(client *Client) {
	request := newRequest(client.APIKey)
	resp, body , err := request.Get(API_URL + "v1/app/list").End()
	fmt.Println(resp, body, err)
}

func ListBackups(client *Client) {
	request := newRequest(client.APIKey)
	resp, body , err := request.Get(API_URL + "v1/backup/list").End()
	fmt.Println(resp, body, err)
}

func ListOS(client *Client) {
	request := newRequest(client.APIKey)
	resp, body , err := request.Get(API_URL + "v1/os/list").End()
	fmt.Println(resp, body, err)
}