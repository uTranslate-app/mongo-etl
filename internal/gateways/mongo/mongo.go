package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/uTranslate-app/uTranslate-api/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type langsResult struct {
	Id    primitive.ObjectID `bson:"_id"`
	Langs []string           `bson:"uniqueValues"`
}

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

	coll := client.Database(configs.Cfg.DbName).Collection(configs.Cfg.SentColl)

	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Documents inserted: %v, in collection %v\n", len(result.InsertedIDs), coll.Name())
}

func (m MongoDb) GetMongoLangs(lang string) []string {
	client := m.ConnectMongo()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err.Error())
		}
	}()
	coll := client.Database(configs.Cfg.DbName).Collection("sentences")
	filter := bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"$or",
						bson.A{
							bson.D{{"sent_a.lang", lang}},
							bson.D{{"sent_b.lang", lang}},
						},
					},
				},
			},
		},
		bson.D{{"$project", bson.D{{"lang", "$sent_a.lang"}}}},
		bson.D{
			{"$unionWith",
				bson.D{
					{"coll", "sentences"},
					{"pipeline",
						bson.A{
							bson.D{{"$project", bson.D{{"lang", "$sent_b.lang"}}}},
						},
					},
				},
			},
		},
		bson.D{{"$match", bson.D{{"lang", bson.D{{"$not", primitive.Regex{Pattern: lang}}}}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", primitive.Null{}},
					{"uniqueValues", bson.D{{"$addToSet", "$lang"}}},
				},
			},
		},
	}
	cursor, err := coll.Aggregate(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []langsResult
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results[0].Langs
}
