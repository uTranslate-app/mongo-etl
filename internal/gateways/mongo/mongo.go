package mongo

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/uTranslate-app/uTranslate-api/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	Uri string
}

func (m MongoDb) ConnectMongo() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(configs.Cfg.MongoUri))
	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}

func (m MongoDb) InsertSentences(file string, docs []interface{}) {
	client := m.ConnectMongo()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err.Error())
		}
	}()

	collName := strings.Split(file, "/")[0]
	coll := client.Database(configs.Cfg.DbName).Collection(collName)

	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Documents inserted: %v, in collection %v\n", len(result.InsertedIDs), coll.Name())
}
