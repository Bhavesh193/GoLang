package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Bhavesh2001:xyMMdqBtZbzmu34C@cluster0.tcxek77.mongodb.net/"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection


func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connnect to the mongodb
	client , err := mongo.Connect(context.TODO() , clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is empty")
}