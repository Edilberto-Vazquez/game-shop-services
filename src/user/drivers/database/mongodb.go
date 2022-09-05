package database

import (
	"context"
	"log"

	"github.com/Edilberto-Vazquez/game-shop-services/src/user/config"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/domains"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

type mongoUser struct {
	ID        uuid.UUID `bson:"id"`
	UserName  string    `bson:"user_name"`
	Email     string    `bson:"email"`
	CountryId string    `bson:"country_id"`
	Salt      string    `bson:"salt"`
	Hash      string    `bson:"hash"`
}

func NewFromUser(user domains.User) mongoUser {
	return mongoUser{
		ID:        user.GetID(),
		UserName:  user.GetUserName(),
		Email:     user.GetEmail(),
		CountryId: user.GetCountryId(),
		Salt:      user.GetSalt(),
		Hash:      user.GetHash(),
	}
}

func NewMongoRepository(conf config.DBConfig) (*MongoRepository, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.URI))
	if err != nil {
		return nil, err
	}
	db := client.Database(conf.Name)
	coll := db.Collection(conf.Collection)
	return &MongoRepository{db: db, coll: coll}, nil
}

func (mongo *MongoRepository) InsertUser(ctx context.Context, user domains.User) (userId string, err error) {
	doc := NewFromUser(user)
	result, err := mongo.coll.InsertOne(ctx, doc)
	if err != nil {
		log.Printf("Drivers(InsertUser): %v", err)
		return "", err
	}
	userId = result.InsertedID.(primitive.ObjectID).Hex()
	return
}

func (mongo *MongoRepository) FindUser(ctx context.Context, userName string) (user domains.User, err error) {
	filter := bson.D{{Key: "userName", Value: userName}}
	err = mongo.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domains.User{}, err
	}
	return
}

func (mongo *MongoRepository) UpdateUser(ctx context.Context, user domains.User) (err error) {
	filter := bson.D{{Key: "_id", Value: user.GetID()}}
	doc := NewFromUser(user)
	_, err = mongo.coll.UpdateOne(ctx, filter, doc)
	return
}

func (mongo *MongoRepository) DeleteUser(ctx context.Context, userID uuid.UUID) (err error) {
	if err != nil {
		log.Printf("Drivers(Deleteuser): %v", err)
		return err
	}
	doc := bson.D{{Key: "_id", Value: userID}}
	_, err = mongo.coll.DeleteOne(ctx, doc)
	return
}
