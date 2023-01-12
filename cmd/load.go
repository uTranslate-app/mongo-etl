package etl

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo(mongodb_uri string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodb_uri))
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
