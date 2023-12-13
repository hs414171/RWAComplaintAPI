package models

type Complaint struct {
	HouseNo   int32  `bson:"house_no"`
	Name      string `bson:"name"`
	Complaint string `bson:"complaint"`
	Type      string `bson:"type"`
}

const (
	TypePlumber     = "Plumber"
	TypeElectrician = "Electrician"
)
