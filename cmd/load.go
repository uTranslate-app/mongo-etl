package etl

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}

func insertSentences(coll *mongo.Collection, db string, docs []interface{}) {
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Documents inserted: %v\n", len(result.InsertedIDs))
}
