package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}

func insertSentences(client *mongo.Client, db string, col string, docs []interface{}) {
	coll := client.Database(db).Collection(col)

	result, _ := coll.InsertMany(context.TODO(), docs)
	fmt.Printf("Documents inserted: %v\n", len(result.InsertedIDs))
}
