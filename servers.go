package govultr

import (
	"github.com/parnurzeal/gorequest"
	"strconv"
	"github.com/fatih/structs"
)

func ChangeApp(client *Client, sub int, app int) (bool, []error) {
	form := AppChangeForm{ Sub: sub, App: app }
	mapForm := structs.Map(form)
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/app_change").Set("API-Key", client.APIKey).
		SendMap(mapForm).End()
	return NoResponse(res, errs)
}

func ChangeAppList(client *Client, sub int) {
	gorequest.New().Get(API_URL + "/v1/server/app_change_list").Set("API-Key", client.APIKey).
		Query("SUBID=" + strconv.Itoa(sub)).End()
}

func DisableBackups(client *Client, sub int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/backup_disable").Set("API-Key", client.APIKey).
		SendString("SUBID=" + strconv.Itoa(sub)).End()
	return NoResponse(res, errs)
}

func EnableBackups(client *Client, sub int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/backup_enable").Set("API-Key", client.APIKey).
		SendString("SUBID=" + strconv.Itoa(sub)).End()
	return NoResponse(res, errs)
}

func GetBackupSchedule(client *Client, sub int) {
	gorequest.New().Post(API_URL + "/v1/server/backup_disable").Set("API-Key", client.APIKey).
		SendString("SUBID=" + strconv.Itoa(sub)).End()
}

func SetBackupSchedule(client *Client, sub int, cron_type string, hour int, dow int, dom int) (bool, []error) {
	mapForm := structs.Map(SetBackupScheduleForm{Sub:sub, Cron:cron_type, Hour:hour, DayOfWeek: dow, DayOfMonth: dom})
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/backup_enable").Set("API-Key", client.APIKey).
		SendMap(mapForm).End()
	return NoResponse(res, errs)
}

func GetBandwidth(client *Client, sub int) {
	gorequest.New().Get(API_URL + "/v1/server/bandwidth").Set("API-Key", client.APIKey).
		Query("SUBID=" + strconv.Itoa(sub)).End()
}

func CreateServer(client *Client, form *CreateServerForm) {
	gorequest.New().Post(API_URL + "/v1/server/create").Set("API-Key", client.APIKey).
		SendMap(form).End()
}

func CreateIPv4(client *Client, form *CreateServerIPv4Form) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/create_ipv4").Set("API-Key", client.APIKey).
		SendMap(structs.Map(form)).End()
	return NoResponse(res, errs)
}

func DestroyServer(client *Client, sub int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/destroy").Set("API-Key", client.APIKey).
		SendString("SUBID=" + strconv.Itoa(sub)).End()
	return NoResponse(res, errs)
}

func DestroyIPv4(client *Client, sub int, ipv4 string) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/destroy_ipv4").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func SetGroupFirewall(client *Client, sub int, firewallgroup int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/firewall_group_set").Set("API-Key", client.APIKey).
		End()
	return NoResponse(res, errs)
}

func GetAppInfo(client *Client, sub int) {
	gorequest.New().Get(API_URL + "/v1/server/get_app_info").Set("API-Key", client.APIKey).End()
}

func GetUserData(client *Client, sub int) {
	gorequest.New().Get(API_URL + "/v1/server/get_user_data").Set("API-Key", client.APIKey).End()
}

func HaltServer(client *Client, sub int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/halt").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func EnableIPv6(client *Client, sub int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/ipv6_enable").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func AttachISO(client *Client, sub int, iso int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/iso_attach").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func DetachISO(client *Client, sub int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/iso_detach").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func ISOStatus(client *Client, sub int) {
	gorequest.New().Get(API_URL + "/v1/server/iso_status").Set("API-Key", client.APIKey).End()
}

func SetLabel(client *Client, sub int, label string) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/label_set").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func ListServers(client *Client) (map[string]Server, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "/v1/server/list").Set("API-Key", client.APIKey).End()
	var servers map[string]Server
	errs2 := UnmarshalJson(body, errs, &servers)
	return servers, errs2
}

func ListIPv4(client *Client, sub int) {
	gorequest.New().Get(API_URL + "/v1/server/list_ipv4").Set("API-Key", client.APIKey).End()
}

func ListIPv6(client *Client, sub int) {
	gorequest.New().Get(API_URL + "/v1/server/list_ipv6").Set("API-Key", client.APIKey).End()
}

func ListNeighbors(client *Client, sub int) {
	gorequest.New().Get(API_URL + "/v1/server/neighbors").Set("API-Key", client.APIKey).End()
}

func ChangeOS(client *Client, sub int, os int) {
	gorequest.New().Get(API_URL + "/v1/server/os_change").Set("API-Key", client.APIKey).End()
}

func ChangeOSList(client *Client, sub int) {
	gorequest.New().Post(API_URL + "/v1/server/os_change_list").Set("API-Key", client.APIKey).End()
}

func EnablePrivateNetwork(client *Client, sub int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/private_network_enable").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func RebootServer(client *Client, sub int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/reboot").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func ReinstallServer(client *Client, sub int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/reinstall").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func RestoreBackup(client *Client, sub int, backup string) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/restore_backup").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func RestoreSnapshot(client *Client, sub int, snapshot string) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/restore_snapshot").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func ReverseDefaultIPv4(client *Client, sub int, ipv4 string) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/reverse_default_ipv4").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func ReverseDeleteIPv6(client *Client, sub int, ipv6 string) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/reverse_delete_ipv6").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func ReverseListIPv6(client *Client, sub int) {
	gorequest.New().Get(API_URL + "/v1/server/reverse_list_ipv6").Set("API-Key", client.APIKey).End()
}

func ReverseSetIPv4(client *Client, sub int, ipv4 string, entry string) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/reverse_set_ipv4").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func ReverseSetIPv6(client *Client, sub int, ipv6 string, entry string) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/reverse_set_ipv6").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func SetUserData(client *Client, sub int, userdata string) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/set_user_data").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func StartServer(client *Client, sub int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/start").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func SetTag(client *Client, sub int, tag string) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/tag_set").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func UpgradePlan(client *Client, sub int, vsplan int) (bool, []error) {
	res, _, errs := gorequest.New().Post(API_URL + "/v1/server/upgrade_plan").Set("API-Key", client.APIKey).End()
	return NoResponse(res, errs)
}

func UpgradePlanList(client *Client, sub int) ([]int, []error) {
	_, body, errs := gorequest.New().Get(API_URL + "/v1/server/upgrade_plan_list").Set("API-Key", client.APIKey).
		SendString("SUBID=" + strconv.Itoa(sub)).End()
	var integers []int
	errs2 := UnmarshalJson(body, errs, &integers)
	return integers, errs2
}