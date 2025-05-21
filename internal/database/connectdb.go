package database

import (
	"context"
	"fmt"
	"time"
	"to_do_list/internal/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {

	var err error
	var mongoClient *mongo.Client
	ctx := context.Background()
	ctx1, cancel := context.WithTimeout(ctx, 45*time.Second)
	defer cancel()
	mongoClient, err = mongo.Connect(ctx1, options.Client().ApplyURI(configs.LoadDbUrl()))
	if err != nil {
		fmt.Println("Mongo DB not connected : ", err)
	}
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := mongoClient.Database("golangdb").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return mongoClient
}

// func Disconnect(client *mongo.Client) {
// 	client.Disconnect()
// }

var DB *mongo.Client = Connect()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangdb").Collection(collectionName)
	return collection
}
