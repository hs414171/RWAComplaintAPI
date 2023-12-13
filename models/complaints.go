package models
import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Complaint struct {
	HouseNo   int    `json:"houseno"`
	Name      string `json:"name"`
	Complaint string `json:"complaint"`
	Type      string `json:"type"`
	AllotedTo string `json:"allotedto"`
	CaseID    primitive.ObjectID  `json:"caseid"`
}

