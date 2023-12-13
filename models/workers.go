package models

type Worker struct {
	Name   string  `bson:"name"`
	Available  bool `bson:"available"`
	Expertise string `bson:"expertise"`

}


