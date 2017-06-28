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
	SnapShot string `json:"SNAPSHOTID,omitempty"`
	EnableIPv6 string `json:"enable_ipv6,omitempty"`
	EnablePrivateNetwork string `json:"enable_private_network,omitempty"`
	Label string `json:"label,omitempty"`
	SshKey string `json:"SSHKEYID,omitempty"`
	AutoBackups string `json:"auto_backups,omitempty"`
	App int `json:"APPID,omitempty"`
	UserData string `json:"userdata,omitempty"`
	NotifyActivate string `json:"notify_activate,omitempty"`
	DDosProtection bool `json:"ddos_protection,omitempty"`
	ReservedIPv4 string `json:"reserved_ip_v4,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Tag string `json:"tag,omitempty"`
	FirewallGroupId string `json:"FIREWALLGROUPID,omitempty"`
}

type CreateServerIPv4Form struct {
	Sub int `json:"SUBID"`
	Reboot bool `json:"reboot,omitempty"`
}

type SubForm struct {
	Sub int `json:"SUBID"`
}

type SetBackupScheduleForm struct {
	Sub int `json:"SUBID"`
	Cron string `json:"cron_type"`
	Hour int `json:"hour,omitempty"`
	DayOfWeek int `json:"dow,omitempty"`
	DayOfMonth int `json:"dom,omitempty"`
}