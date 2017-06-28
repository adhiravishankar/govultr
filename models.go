package govultr

import "time"

type Account struct {
	Balance string `json:"balance"`
	PendingCharges string `json:"pending_charges"`
	LastPaymentDate string `json:"last_payment_date"`
	LastPaymentAmount string `json:"last_payment_amount"`
}

type App struct {
	Id string `json:"APPID"`
	Name string `json:"name"`
	ShortName string `json:"short_name"`
	DeployName string `json:"deploy_name"`
	Surcharge int `json:"surcharge"`
}

type Auth struct {
	Acls []string `json:"acls"`
	Email string `json:"email"`
	Name string `json:"name"`
}

type Backup struct {
	Id string `json:"BACKUPID"`
	DateCreated time.Time `json:"date_created"`
	Description string `json:"description"`
	Size int `json:"size"`
	Status string `json:"status"`
}

type BackupSchedule struct {
	Enabled bool `json:"enabled"`
	CronType string `json:"cron_type"`
	NextScheduledTimeUtc string `json:"next_scheduled_time_utc"`
	Hour int `json:"hour"`
	Dow int `json:"dow"`
	Dom int `json:"dom"`
}

type Bandwidth struct {
	IncomingBytes [] BandwidthItem `json:"incoming_bytes"`
	OutgoingBytes [] BandwidthItem `json:"outgoing_bytes"`
}

type BandwidthItem struct {
	Date string `json:"0"`
	Bytes string `json:"1"`
}

type BlockSimple struct {
	UserId string `json:"USERID"`
	APIKey string `json:"api_key"`
}

type Block struct {
	Sub int `json:"SUBID"`
	DateCreated string `json:"date_created"`
	CostPerMonth int `json:"cost_per_month"`
	Status string `json:"status"`
	SizeGb int `json:"size_gb"`
	Datacenter int `json:"DCID"`
	AttachedToSUBID int `json:"attached_to_SUBID"`
	Label string `json:"label"`
}

type DnsDomain struct {
	Domain string `json:"domain"`
	DateCreated string `json:"date_created"`
}

type DnsRecords struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Data string `json:"data"`
	Priority int `json:"priority"`
	Record int `json:"RECORDID"`
	TTL int `json:"ttl"`
}

type Firewall struct {
	Id string `json:"FIREWALLGROUPID"`
	Description string `json:"description"`
	DateCreated time.Time `json:"date_created"`
	DateModified time.Time `json:"date_modified"`
	InstanceCount int `json:"instance_count"`
	RuleCount int `json:"rule_count"`
	MaxRuleCount int `json:"max_rule_count"`
}

type FirewallRule struct {
	Id int `json:"rulenumber"`
	Action string `json:"action"`
	Protocol string `json:"protocol"`
	Port string `json:"port"`
	Subnet string `json:"subnet"`
	SubnetSize int `json:"subnet_size"`
}

type ISO struct {
	Id int `json:"ISOID"`
	DateCreated string `json:"date_created"`
	Filename string `json:"filename"`
	Size int `json:"size"`
	Md5Sum string `json:"md5sum"`
	Status string `json:"status"`
}

type OS struct {
	Id int `json:"OSID"`
	Name string `json:"name"`
	Arch string `json:"arch"`
	Family string `json:"family"`
	Windows bool `json:"windows"`
}

type Plan struct {
	Id string `json:"VPSPLANID"`
	Name string `json:"name"`
	VcpuCount string `json:"vcpu_count"`
	RAM string `json:"ram"`
	Disk string `json:"disk"`
	Bandwidth string `json:"bandwidth"`
	PricePerMonth string `json:"price_per_month"`
	Windows bool `json:"windows"`
	PlanType string `json:"plan_type"`
	AvailableLocations []int `json:"available_locations"`
}

type Region struct {
	Datacenter string `json:"DCID"`
	Name string `json:"name"`
	Country string `json:"country"`
	Continent string `json:"continent"`
	State string `json:"state"`
	DdosProtection bool `json:"ddos_protection"`
	BlockStorage bool `json:"block_storage"`
	Regioncode string `json:"regioncode"`
}

type ReservedIP struct {
	Sub int `json:"SUBID"`
	Datacenter int `json:"DCID"`
	IPType string `json:"ip_type"`
	Subnet string `json:"subnet"`
	SubnetSize int `json:"subnet_size"`
	Label string `json:"label"`
	AttachedSub int `json:"attached_SUBID"`
}

type SnapShot struct {
	Id string `json:"SNAPSHOTID"`
	DateCreated string `json:"date_created"`
	Description string `json:"description"`
	Size string `json:"size"`
	Status string `json:"status"`
}

type SshKey struct {
	Id string `json:"SSHKEYID"`
	DateCreated interface{} `json:"date_created"`
	Name string `json:"name"`
	SSHKey string `json:"ssh_key"`
}

type StartupScript struct {
	Id string `json:"SCRIPTID"`
	DateCreated string `json:"date_created"`
	DateModified string `json:"date_modified"`
	Name string `json:"name"`
	Type string `json:"type"`
	Script string `json:"script"`
}

type Server struct {
	SubId string `json:"SUBID"`
	OS string `json:"os"`
	RAM string `json:"ram"`
	Disk string `json:"disk"`
	MainIP string `json:"main_ip"`
	VcpuCount string `json:"vcpu_count"`
	Location string `json:"location"`
	DCID string `json:"DCID"`
	DefaultPassword string `json:"default_password"`
	DateCreated string `json:"date_created"`
	PendingCharges string `json:"pending_charges"`
	Status string `json:"status"`
	CostPerMonth string `json:"cost_per_month"`
	CurrentBandwidthGb float64 `json:"current_bandwidth_gb"`
	AllowedBandwidthGb string `json:"allowed_bandwidth_gb"`
	NetmaskV4 string `json:"netmask_v4"`
	GatewayV4 string `json:"gateway_v4"`
	PowerStatus string `json:"power_status"`
	ServerState string `json:"server_state"`
	VPSPLANID string `json:"VPSPLANID"`
	V6Network string `json:"v6_network"`
	V6MainIP string `json:"v6_main_ip"`
	V6NetworkSize string `json:"v6_network_size"`
	V6Networks []V6Network `json:"v6_networks"`
	Label string `json:"label"`
	InternalIP string `json:"internal_ip"`
	KvmURL string `json:"kvm_url"`
	AutoBackups string `json:"auto_backups"`
	Tag string `json:"tag"`
	OSId string `json:"OSID"`
	AppId string `json:"APPID"`
	FirewallGroupId string `json:"FIREWALLGROUPID"`
}

type V6Network struct {
	V6Network string `json:"v6_network"`
	V6MainIP string `json:"v6_main_ip"`
	V6NetworkSize string `json:"v6_network_size"`
}