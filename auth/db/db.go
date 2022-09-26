package db

// rename to mongoClient

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURL = "mongodb://localhost:27017"

var Client, Ctx = MongoClient()

func MongoClient() (*mongo.Client, context.Context) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "root",
		Password: "pass",
	})

	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("err connecting mongo", err)
	}
	return client, ctx
}

func Health() bool {
	err := Client.Ping(Ctx, nil)
	if err != nil {
		log.Println("error connecting to mongo")
		return false
	}
	return true
}

func init() {
	initDb()
}

func initDb() {
	log.Println("init db")

	var db = Client.Database("USERS_DB")

	db.Collection("users").Indexes().CreateOne(Ctx, mongo.IndexModel{
		Keys: bson.M{
			"userName": 1,
		},
		Options: options.Index().SetUnique(true),
	})

}
