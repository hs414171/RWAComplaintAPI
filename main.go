package main

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "gofr.dev/pkg/gofr"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main() {
    
    if err := godotenv.Load("configs/.env"); err != nil {
        log.Fatal("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGO_URI")

    opts := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), opts)
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
    app := gofr.New()
    app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
        return "Hello World!", nil
    })

    
    app.Start()
}
