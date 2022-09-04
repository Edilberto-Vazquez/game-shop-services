package database

import (
	"context"
	"log"

	"github.com/Edilberto-Vazquez/game-shop-services/src/user/config"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (mongo *MongoDB) InsertUser(ctx context.Context, user *models.User) (userId string, err error) {
	doc := bson.D{
		{Key: "user_name", Value: user.UserName},
		{Key: "email", Value: user.Email},
		{Key: "country_id", Value: user.CountryId},
		{Key: "salt", Value: user.Salt},
		{Key: "hash", Value: user.Hash},
	}
	result, err := mongo.coll.InsertOne(ctx, doc)
	if err != nil {
		log.Printf("Drivers(InsertUser): %v", err)
		return "", err
	}
	userId = result.InsertedID.(primitive.ObjectID).Hex()
	return
}

func (mongo *MongoDB) FindUser(ctx context.Context, userId string) (user *models.User, err error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Printf("Drivers(FindUser): %v", err)
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: id}}
	err = mongo.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return
}

func (mongo *MongoDB) UpdateUser(ctx context.Context, user *models.User) (err error) {
	id, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		log.Printf("Drivers(Updateuser): %v", err)
		return err
	}
	filter := bson.D{{Key: "_id", Value: id}}
	doc := bson.D{
		{Key: "user_name", Value: user.UserName},
		{Key: "email", Value: user.Email},
		{Key: "country_id", Value: user.CountryId},
		{Key: "salt", Value: user.Salt},
		{Key: "hash", Value: user.Hash},
	}
	_, err = mongo.coll.UpdateOne(ctx, filter, doc)
	return
}

func (mongo *MongoDB) DeleteUser(ctx context.Context, userId string) (err error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Printf("Drivers(Deleteuser): %v", err)
		return err
	}
	doc := bson.D{{Key: "_id", Value: id}}
	_, err = mongo.coll.DeleteOne(ctx, doc)
	return
}
