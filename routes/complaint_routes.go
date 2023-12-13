package routes

import (
	"log"

	"github.com/hs414171/AVRWA_COMPLAINT/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gofr.dev/pkg/gofr"
)


type CaseID struct {
	CaseId primitive.ObjectID `json:"case_id"`
}

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
	complaint.CaseID = primitive.NewObjectID()
	collection := client.Database("RWA").Collection("Complaints")

	_, err := collection.InsertOne(ctx, complaint)
	if err != nil {
		return nil, err
	}

	return complaint, nil

}

func DeleteComplaintByCaseID(ctx *gofr.Context, client *mongo.Client) (interface{}, error) {
	var cAseId = ctx.PathParam("case_id")
	objID, err := primitive.ObjectIDFromHex(cAseId)
	if err != nil {
		return primitive.NilObjectID, err
	}
	log.Println(objID)
	collection := client.Database("RWA").Collection("Complaints")

	filter := bson.M{"caseid": objID}
	log.Println(filter)

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}


