package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hs414171/AVRWA_COMPLAINT/routes"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello World!", nil
	})

	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")

	opts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), opts)
	log.Println("server", client)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	app.GET("/getcomplaints", func(ctx *gofr.Context) (interface{}, error) {
		return routes.GetAllComplaints(ctx, client)
	})

	app.GET("/complaint/{case_id}", func(ctx *gofr.Context) (interface{}, error) {
		return routes.FindComplaintsByID(ctx, client)
	})

	app.POST("/addcomplaints", func(ctx *gofr.Context) (interface{}, error) {
		return routes.HandleComplaints(ctx, client)
	})

	app.DELETE("/delcomp/{case_id}", func(ctx *gofr.Context) (interface{}, error) {
		return routes.DeleteComplaintByCaseID(ctx, client)
	})

	app.PATCH("/updatecomp/{case_id}", func(ctx *gofr.Context) (interface{}, error) {
		return routes.UpdateComplaintsByCaseID(ctx, client)
	})

	app.GET("/getworkers", func(ctx *gofr.Context) (interface{}, error) {
		return routes.GetAllWorkers(ctx, client)
	})

	app.GET("/worker/{emp_id}", func(ctx *gofr.Context) (interface{}, error) {
		return routes.FindWorkersByID(ctx, client)
	})

	app.POST("/addworkers", func(ctx *gofr.Context) (interface{}, error) {
		return routes.HandleWorkers(ctx, client)
	})

	app.DELETE("/delworker/{emp_id}", func(ctx *gofr.Context) (interface{}, error) {
		return routes.DeleteWorkerByCaseID(ctx, client)
	})

	app.PATCH("/updateworker/{emp_id}", func(ctx *gofr.Context) (interface{}, error) {
		return routes.UpdateWorkerByCaseID(ctx, client)
	})

	app.Start()
}
