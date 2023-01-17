package db

import (
	"context"
	"log"

	"github.com/uTranslate-app/uTranslate-api/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(configs.Cfg.MongoUri))
	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}
