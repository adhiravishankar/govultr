package govultr

type OS struct {
	Id int `json:"OSID"`
	Name string `json:"name"`
	Arch string `json:"arch"`
	Family string `json:"family"`
	Windows bool `json:"windows"`
}