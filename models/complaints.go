package models
import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Complaint struct {
	HouseNo   int    `json:"house_no"`
	Name      string `json:"name"`
	Complaint string `json:"complaint"`
	Type      string `json:"type"`
	AllotedTo string `json:"alloted_to"`
	CaseID    primitive.ObjectID  `json:"case_id"`
}

