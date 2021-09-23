package db

import (
	"context"
	"log"

	middlewares "github.com/arykalin/kogda-sobitie-backend/handlers"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Dbconnect -> connects mongo
func Dbconnect() *mongo.Client {
	url := middlewares.DotEnvVariable("MONGO_URL")
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("⛒ Connection Failed to Database on url %s reason: %s", url, err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("⛒ Connection Failed to Database on url %s reason: %s", url, err)
	}
	color.Green("⛁ Connected to Database")
	return client
}
