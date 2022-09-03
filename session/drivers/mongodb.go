package drivers

import (
	"context"
	"log"

	"github.com/Edilberto-Vazquez/game-shop-services/session/config"
	"github.com/Edilberto-Vazquez/game-shop-services/session/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	coll *mongo.Collection
}

func NewMongoDB(conf config.DBConfig) *MongoDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.URI))
	collection := client.Database(conf.Name).Collection(conf.Collection)

	if err != nil {
		log.Fatal(err)
	}

	return &MongoDB{collection}
}

func (mongo *MongoDB) InsertUser(ctx context.Context, user *models.User) (err error) {
	doc := bson.D{
		{Key: "user_name", Value: user.UserName},
		{Key: "email", Value: user.Email},
		{Key: "country_id", Value: user.CountryId},
		{Key: "salt", Value: user.Salt},
		{Key: "hash", Value: user.Hash},
	}

	_, err = mongo.coll.InsertOne(ctx, doc)

	return
}
