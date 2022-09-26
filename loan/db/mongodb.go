package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURL = "mongodb://localhost:1244"

var Client, Ctx = MongoClient()

func init() {
	initDB()
}

func MongoClient() (*mongo.Client, context.Context) {

	clientOptions := options.Client().ApplyURI(mongoURL)
	// clientOptions.SetAuth(options.Credential{
	// 	Username: "",
	// 	Password: "",
	// })

	ctx := context.Background()
	Client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("err connecting to mongoDB")
	}
	return Client, ctx
}

func initDB() {
	log.Println("initing DB...")
	var db = Client.Database("LOANS_DB")
	db.Collection("loans").Indexes().CreateOne(Ctx, mongo.IndexModel{
		Keys: bson.M{
			"userName": 1,
		},
		Options: options.Index().SetUnique(true),
	})
}
