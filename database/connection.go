package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const (
	DbName = "translator"
)

func NewMongo() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:password@127.0.0.1:27017/translator?authSource=admin"))
	if err != nil {
		log.Fatal("Error getting client: ", err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting: ", err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error pinging: ", err.Error())
	}
	return client
}
