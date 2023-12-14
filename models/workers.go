package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Worker struct {
	Name      string `json:"name"`
	Available *bool   `json:"available"`
	Expertise string `json:"expertise"`
	EmpID     primitive.ObjectID `json:"empid"`
	AssignedCases []primitive.ObjectID `json:"assignedcases"`
	AddCaseID     primitive.ObjectID    `json:"addcaseid"`
	RemoveCaseID  primitive.ObjectID    `json:"removecaseid"`
}
