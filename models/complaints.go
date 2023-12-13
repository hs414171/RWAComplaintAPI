package models

type Complaint struct {
	HouseNo   int    `json:"house_no"`
	Name      string `json:"name"`
	Complaint string `json:"complaint"`
	Type      string `json:"type"`
}

const (
	TypePlumber     = "Plumber"
	TypeElectrician = "Electrician"
)
