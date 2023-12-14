package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Worker struct {
	Name      string `bson:"name"`
	Available *bool   `bson:"available"`
	Expertise string `bson:"expertise"`
	EmpID     primitive.ObjectID `json:"empid"`
}
