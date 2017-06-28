package govultr

type AppChangeForm struct {
	Sub int `json:"SUBID"`
	App int `json:"APPID"`
}

type CreateServerForm struct {
	Datacenter int `json:"DCID"`
	VPSPlan int `json:"VPSPLANID"`
	OS int `json:"OSID"`
	IPXEChainURL string `json:"ipxe_chain_url,omitempty"`
	ISO string `json:"ISOID,omitempty"`
	Script int `json:"SCRIPTID,omitempty"`
}

type SetBackupScheduleForm struct {
	Sub int `json:"SUBID"`
	Cron string `json:"cron_type"`
	Hour int `json:"hour,omitempty"`
	DayOfWeek int `json:"dow,omitempty"`
	DayOfMonth int `json:"dom,omitempty"`
}