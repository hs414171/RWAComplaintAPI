package routes

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gofr.dev/pkg/gofr"
)

type Complaint struct {
	Houseno   int32  `bson:"house_no"`
	Name      string `bson:"name"`
	Complaint string `bson:"complaint"`
	Type      string `bson:"type"`
}

func GetAllComplaints(ctx *gofr.Context, client *mongo.Client) (interface{}, error) {

	collection := client.Database("RWA").Collection("Complaints")

	filter := bson.D{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var complaints []Complaint
	for cursor.Next(ctx) {
		var complaint Complaint
		if err := cursor.Decode(&complaint); err != nil {
			return nil, err
		}
		complaints = append(complaints, complaint)
	}

	return complaints, nil

}
