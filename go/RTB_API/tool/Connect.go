package tool

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//连接mongodb
func ConnectDB() *mongo.Client {
	//建立连接
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var ctx = context.TODO()
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
func ConnectCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	//连接到指定db的collection
	collection := client.Database("oyc_task").Collection(collectionName)
	return collection
}
