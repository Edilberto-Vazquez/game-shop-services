package database

import (
	"context"
	"time"

	"github.com/Edilberto-Vazquez/game-shop-services/src/user/config"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/domains"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

type mongoDBUser struct {
	id        primitive.ObjectID `bson:"_id"`
	userName  string             `bson:"user_name"`
	email     string             `bson:"email"`
	countryId string             `bson:"country_id"`
	password  string             `bson:"password"`
	createdAt time.Time          `bson:"created_at"`
	updatedAt time.Time          `bson:"updated_at"`
	deletedAt time.Time          `bson:"deleted_at,omitempty"`
}

func NewFromUser(user domains.User) mongoDBUser {
	return mongoDBUser{
		id:        primitive.NewObjectID(),
		userName:  user.GetUserName(),
		email:     user.GetEmail(),
		countryId: user.GetCountryId(),
		password:  user.GetPassword(),
		createdAt: time.Now(),
	}
}

func NewMongoDBRepository(conf config.DBConfig) (*MongoDBRepository, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.URI))
	if err != nil {
		return nil, err
	}
	db := client.Database(conf.Name)
	coll := db.Collection(conf.Collection)
	return &MongoDBRepository{db: db, coll: coll}, nil
}

func (mongo *MongoDBRepository) InsertUser(ctx context.Context, user domains.User) (id string, err error) {
	doc := NewFromUser(user)
	result, err := mongo.coll.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}
	id = result.InsertedID.(primitive.ObjectID).Hex()
	return
}

func (mongo *MongoDBRepository) FindUserById(ctx context.Context, id primitive.ObjectID) (user domains.User, err error) {
	filter := bson.M{"_id": id, "deleted_at": bson.M{"$exist": false}}
	err = mongo.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domains.User{}, err
	}
	return
}

func (mongo *MongoDBRepository) FindUserByEmail(ctx context.Context, email string) (user domains.User, err error) {
	filter := bson.M{"email": email, "deleted_at": bson.M{"$exist": false}}
	err = mongo.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domains.User{}, err
	}
	return
}

func (mongo *MongoDBRepository) UpdateUser(ctx context.Context, user domains.User) (err error) {
	filter := bson.M{"_id": user.GetID(), "deleted_at": bson.M{"$exist": false}}
	doc := NewFromUser(user)
	doc.updatedAt = time.Now()
	_, err = mongo.coll.UpdateOne(ctx, filter, doc)
	return
}

func (mongo *MongoDBRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) (err error) {
	filter := bson.M{"_id": id, "deleted_at": bson.M{"$exist": false}}
	updater := bson.D{{Key: "$set", Value: bson.D{{Key: "deleted_at", Value: time.Now()}}}}
	_, err = mongo.coll.UpdateOne(ctx, filter, updater)
	return
}
