package routes

import (
	"log"
	// "encoding/json"
	"github.com/hs414171/AVRWA_COMPLAINT/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gofr.dev/pkg/gofr"
)

// type Complaint struct {
// 	Houseno   string  `bson:"house_no"`
// 	Name      string `bson:"name"`
// 	Complaint string `bson:"complaint"`
// 	Type      string `bson:"type"`
// }

func GetAllComplaints(ctx *gofr.Context, client *mongo.Client) (interface{}, error) {

	collection := client.Database("RWA").Collection("Complaints")

	filter := bson.D{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var complaints []models.Complaint
	for cursor.Next(ctx) {
		var complaint models.Complaint
		if err := cursor.Decode(&complaint); err != nil {
			return nil, err
		}
		complaints = append(complaints, complaint)
	}

	return complaints, nil

}

func HandleComplaints(ctx *gofr.Context, client *mongo.Client) (interface{}, error) {

	var complaint models.Complaint
	ctx.Bind(&complaint)
	log.Println(complaint)
	collection := client.Database("RWA").Collection("Complaints")

	_, err := collection.InsertOne(ctx, complaint)
	if err != nil {
		return nil, err
	}

	return complaint, nil

}
