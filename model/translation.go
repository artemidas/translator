package model

import (
	"context"
	"github.com/artemidas/translator/database"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Translation struct {
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (t *Translation) Insert(db *mongo.Client, collection string) error {
	c := db.Database(database.DbName).Collection(collection)
	_, err := c.InsertOne(context.TODO(), t)
	defer db.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}
