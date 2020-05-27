package data

import (
	"context"
	"fmt"
	"log"

	"github.com/Comment-API/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection of comments
var Collection *mongo.Collection

// DatabaseConnection :
func DatabaseConnection() {
	fmt.Println("Connected to Database!")
	clientOptions := options.Client().ApplyURI(config.C.Database.Addr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database(config.C.Database.DBName).Collection("comments")

}
