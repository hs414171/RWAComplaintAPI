package routes

import (
	"log"

	"github.com/hs414171/AVRWA_COMPLAINT/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gofr.dev/pkg/gofr"
)

func GetAllWorkers(ctx *gofr.Context, client *mongo.Client) (interface{}, error) {

	collection := client.Database("RWA").Collection("Workers")

	filter := bson.D{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var workers []models.Worker
	for cursor.Next(ctx) {
		var worker models.Worker
		if err := cursor.Decode(&worker); err != nil {
			return nil, err
		}
		workers = append(workers, worker)
	}

	return workers, nil

}

func HandleWorkers(ctx *gofr.Context, client *mongo.Client) (interface{}, error) {

	var worker models.Worker
	ctx.Bind(&worker)
	worker.EmpID = primitive.NewObjectID()
	collection := client.Database("RWA").Collection("Workers")

	_, err := collection.InsertOne(ctx, worker)
	if err != nil {
		return nil, err
	}

	return worker, nil

}

func DeleteWorkerByCaseID(ctx *gofr.Context, client *mongo.Client) (interface{}, error) {
	var empId = ctx.PathParam("emp_id")
	objID, err := primitive.ObjectIDFromHex(empId)
	if err != nil {
		return primitive.NilObjectID, err
	}
	collection := client.Database("RWA").Collection("Workers")

	filter := bson.M{"empid": objID}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}

func UpdateWorkerByCaseID(ctx *gofr.Context, client *mongo.Client) (interface{}, error) {
	var empID = ctx.PathParam("emp_id")
	var updatedFields models.Worker
	ctx.Bind(&updatedFields)
	objID, err := primitive.ObjectIDFromHex(empID)
	if err != nil {
		return primitive.NilObjectID, err
	}
	collection := client.Database("RWA").Collection("Workers")

	filter := bson.M{"empid": objID}
	update := bson.M{"$set": bson.M{}}

	if updatedFields.Name != "" {
		update["$set"].(bson.M)["name"] = updatedFields.Name
	}
	if updatedFields.Available != false {
		update["$set"].(bson.M)["available"] = updatedFields.Available
	}
	if updatedFields.Expertise != "" {
		update["$set"].(bson.M)["expertise"] = updatedFields.Expertise
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return result.ModifiedCount, nil

}