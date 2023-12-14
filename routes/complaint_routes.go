package routes

import (
	"log"
	"sort"

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

	var workers []models.Worker
	filter := bson.M{"expertise": complaint.Type, "available": true}
	collectionWorkers := client.Database("RWA").Collection("Workers")
	cursor, err := collectionWorkers.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &workers)
	if err != nil {
		return nil, err
	}

	var selectedWorker models.Worker
	if len(workers) > 0 {

		sort.Slice(workers, func(i, j int) bool {
			return len(workers[i].AssignedCases) < len(workers[j].AssignedCases)
		})

		selectedWorker = workers[0]

		update := bson.M{
			"$push": bson.M{"assignedcases": complaint.CaseID},
			"$set":  bson.M{"available": len(selectedWorker.AssignedCases) < 4},
		}
		_, err = collectionWorkers.UpdateOne(ctx, bson.M{"empid": selectedWorker.EmpID}, update)
		if err != nil {
			return nil, err
		}

		complaint.AllotedTo = selectedWorker.EmpID
	}

	_, err = collection.InsertOne(ctx, complaint)
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
	collection := client.Database("RWA").Collection("Complaints")

	filter := bson.M{"caseid": objID}
	log.Println(filter)

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}

func UpdateComplaintsByCaseID(ctx *gofr.Context, client *mongo.Client) (interface{}, error) {
	var caseID = ctx.PathParam("case_id")
	var updatedFields models.Complaint

	ctx.Bind(&updatedFields)

	objID, err := primitive.ObjectIDFromHex(caseID)
	if err != nil {
		return primitive.NilObjectID, err
	}

	collection := client.Database("RWA").Collection("Complaints")
	collectionWorkers := client.Database("RWA").Collection("Workers")

	filter := bson.M{"caseid": objID}
	res := collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	var existingComplaint models.Complaint
	if err := res.Decode(&existingComplaint); err != nil {
		return nil, err
	}

	update := bson.M{"$set": bson.M{}}

	if updatedFields.Name != "" {
		update["$set"].(bson.M)["name"] = updatedFields.Name
	}
	if updatedFields.HouseNo != 0 {
		update["$set"].(bson.M)["houseno"] = updatedFields.HouseNo
	}
	if updatedFields.Complaint != "" {
		update["$set"].(bson.M)["complaint"] = updatedFields.Complaint
	}
	if updatedFields.Type != "" {
		update["$set"].(bson.M)["type"] = updatedFields.Type
	}
	if updatedFields.AllotedTo != primitive.NilObjectID {
		update["$set"].(bson.M)["allotedto"] = updatedFields.AllotedTo

		if updatedFields.AllotedTo != existingComplaint.AllotedTo {

			updateNewWorker := bson.M{"$push": bson.M{"assignedcases": objID}}
			_, err := collectionWorkers.UpdateOne(ctx, bson.M{"empid": updatedFields.AllotedTo}, updateNewWorker)
			if err != nil {
				return primitive.NilObjectID, err
			}

			if existingComplaint.AllotedTo != primitive.NilObjectID {
				updatePrevWorker := bson.M{"$pull": bson.M{"assignedcases": objID}}
				_, err = collectionWorkers.UpdateOne(ctx, bson.M{"empid": existingComplaint.AllotedTo}, updatePrevWorker)
				if err != nil {
					return primitive.NilObjectID, err
				}
			}
		}
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result.ModifiedCount, nil
}
