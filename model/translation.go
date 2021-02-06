package model

import (
	"context"
	"errors"
	"github.com/Jeffail/gabs"
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

func (t *Translation) GetLocale(db *mongo.Client, locale string) (string, error) {
	collection := db.Database(database.DbName).Collection(locale)
	defer db.Disconnect(context.Background())

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	var rows []bson.M
	if err = cursor.All(context.Background(), &rows); err != nil {
		log.Fatal(err)
		return "", err
	}
	// Transform to key value object
	translations := gabs.New()
	for _, row := range rows {
		_, err := translations.SetP(row["value"].(string), row["key"].(string))
		if err != nil {
			log.Fatal(err)
			return "", err
		}
	}

	return translations.String(), nil
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

func (t *Translation) Update(db *mongo.Client, locale string, objectID string) error {
	c := db.Database(database.DbName).Collection(locale)
	id, _ := primitive.ObjectIDFromHex(objectID)
	if id.IsZero() {
		return errors.New("record not found")
	}
	_, err := c.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{{"$set", bson.D{{"value", t.Value}}}})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
