package model

import (
	"context"
	"github.com/artemidas/translator/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Translation struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Key       string             `json:"key" bson:"key"`
	Value     string             `json:"value" bson:"value"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (t *Translation) GetLocale(db *mongo.Client, locale string) ([]bson.M, error) {
	collection := db.Database(database.DbName).Collection(locale)
	defer db.Disconnect(context.Background())

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var translations []bson.M
	if err = cursor.All(context.Background(), &translations); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return translations, nil
}

func (t *Translation) Insert(db *mongo.Client, locale string) error {
	c := db.Database(database.DbName).Collection(locale)
	t.CreatedAt = time.Now().UTC()
	t.UpdatedAt = time.Now().UTC()
	_, err := c.InsertOne(context.TODO(), t)
	defer db.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}
