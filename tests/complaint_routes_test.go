package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hs414171/AVRWA_COMPLAINT/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gofr.dev/pkg/gofr"
)


type MockMongoClient struct {
	Complaints map[string]models.Complaint

}

func NewMockMongoClient() *MockMongoClient {
	return &MockMongoClient{
		Complaints: make(map[string]models.Complaint),
	}
}

func (mc *MockMongoClient) InsertComplaint(ctx *gofr.Context, complaint models.Complaint) error {
	mc.Complaints[complaint.CaseID.Hex()] = complaint
	return nil
}

func (mc *MockMongoClient) GetComplaint(ctx *gofr.Context, caseID primitive.ObjectID) (models.Complaint, error) {
	if complaint, found := mc.Complaints[caseID.Hex()]; found {
		return complaint, nil
	}
	return models.Complaint{}, mongo.ErrNoDocuments
}

func (mc *MockMongoClient) GetAllComplaints(ctx *gofr.Context) ([]models.Complaint, error) {
	var complaints []models.Complaint
	for _, complaint := range mc.Complaints {
		complaints = append(complaints, complaint)
	}
	return complaints, nil
}

func (mc *MockMongoClient) DeleteComplaint(ctx *gofr.Context, caseID primitive.ObjectID) error {
	delete(mc.Complaints, caseID.Hex())
	return nil
}

func (mc *MockMongoClient) UpdateComplaint(ctx *gofr.Context, caseID primitive.ObjectID, updatedFields models.Complaint) error {
	if _, found := mc.Complaints[caseID.Hex()]; !found {
		return mongo.ErrNoDocuments
	}

	mc.Complaints[caseID.Hex()] = updatedFields
	return nil
}

type MockGoFrContext struct {
	RequestBody []byte
	Parameters  map[string]string
}

func NewMockGoFrContext(body []byte, params map[string]string) *MockGoFrContext {
	return &MockGoFrContext{
		RequestBody: body,
		Parameters:  params,
	}
}

func (ctx *MockGoFrContext) Bind(obj interface{}) error {
	if len(ctx.RequestBody) == 0 {
		return errors.New("no request body provided")
	}
	err := json.Unmarshal(ctx.RequestBody, obj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal request body: %w", err)
	}
	return nil
}

func (ctx *MockGoFrContext) PathParam(param string) string {
	val, found := ctx.Parameters[param]
	if !found {
		return ""
	}
	return val
}

func FindComplaintsByID(ctx *gofr.Context, client *MockMongoClient) (interface{}, error) {
	var caseID = ctx.PathParam("case_id")
	objID, err := primitive.ObjectIDFromHex(caseID)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return client.GetComplaint(ctx, objID)
}

func GetAllComplaints(ctx *gofr.Context, client *MockMongoClient) (interface{}, error) {
	return client.GetAllComplaints(ctx)
}

func HandleComplaints(ctx *gofr.Context, client *MockMongoClient) (interface{}, error) {
	var complaint models.Complaint
	ctx.Bind(&complaint)
	complaint.CaseID = primitive.NewObjectID()
	return nil, client.InsertComplaint(ctx, complaint)
}

func DeleteComplaintByCaseID(ctx *gofr.Context, client *MockMongoClient) (interface{}, error) {
	var cAseId = ctx.PathParam("case_id")
	objID, err := primitive.ObjectIDFromHex(cAseId)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return nil, client.DeleteComplaint(ctx, objID)
}

func UpdateComplaintsByCaseID(ctx *gofr.Context, client *MockMongoClient) (interface{}, error) {
	var caseID = ctx.PathParam("case_id")
	var updatedFields models.Complaint

	ctx.Bind(&updatedFields)

	objID, err := primitive.ObjectIDFromHex(caseID)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return nil, client.UpdateComplaint(ctx, objID, updatedFields)
}

func main() {
	mongoClient := NewMockMongoClient()
	router := gofr.New()

	router.GET("/complaint/{case_id}", func(ctx *gofr.Context) (interface{}, error) {
		handler, err := FindComplaintsByID(ctx, mongoClient)
		return handler,err

	})
	router.POST("/addcomplaints", func(ctx *gofr.Context) (interface{}, error) {
		handler, err := HandleComplaints(ctx, mongoClient)
		return handler,err
	})

	router.PATCH("/updatecomp/{case_id}", func(ctx *gofr.Context) (interface{}, error) {
		handler, err := UpdateComplaintsByCaseID(ctx, mongoClient)
		return handler,err
	})

	router.DELETE("/delcomp/{case_id}", func(ctx *gofr.Context) (interface{}, error) {
		handler, err := DeleteComplaintByCaseID(ctx, mongoClient)
		return handler,err
	})

}
