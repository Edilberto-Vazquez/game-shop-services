package user

import (
	"context"
	"time"

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
	Id        primitive.ObjectID `bson:"_id"`
	UserName  string             `bson:"user_name"`
	Email     string             `bson:"email"`
	CountryId string             `bson:"country_id"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty"`
}

func NewFromUser(user User) mongoDBUser {
	return mongoDBUser{
		UserName:  user.GetUserName(),
		Email:     user.GetEmail(),
		CountryId: user.GetCountryId(),
		Password:  user.GetPassword(),
	}
}

func NewMongoDBRepository(conf Config) (*MongoDBRepository, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.dbUri))
	if err != nil {
		return nil, err
	}
	db := client.Database(conf.dbName)
	coll := db.Collection(conf.dbCollection)
	return &MongoDBRepository{db: db, coll: coll}, nil
}

func (mongo *MongoDBRepository) InsertUser(ctx context.Context, user User) (err error) {
	doc := NewFromUser(user)
	doc.UpdatedAt = time.Now()
	doc.Id = primitive.NewObjectID()
	result, err := mongo.coll.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	user.SetID(result.InsertedID.(primitive.ObjectID).Hex())
	return
}

func (mongo *MongoDBRepository) FindUserById(ctx context.Context, id primitive.ObjectID) (user User, err error) {
	filter := bson.M{"_id": id, "deleted_at": bson.M{"$exist": false}}
	err = mongo.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return
}

func (mongo *MongoDBRepository) FindUserByEmail(ctx context.Context, email string) (user User, err error) {
	filter := bson.M{"email": email, "deleted_at": bson.M{"$exists": false}}
	var person Person
	err = mongo.coll.FindOne(ctx, filter).Decode(&person)
	if err != nil {
		return user, err
	}
	user = NewUser(&person)
	return
}

func (mongo *MongoDBRepository) UpdateUser(ctx context.Context, user User) (err error) {
	filter := bson.M{"_id": user.GetID(), "deleted_at": bson.M{"$exist": false}}
	doc := NewFromUser(user)
	doc.UpdatedAt = time.Now()
	_, err = mongo.coll.UpdateOne(ctx, filter, doc)
	return err
}

func (mongo *MongoDBRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) (err error) {
	filter := bson.M{"_id": id, "deleted_at": bson.M{"$exist": false}}
	updater := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "deleted_at", Value: time.Now()},
				{Key: "updated_at", Value: time.Now()},
			}},
	}
	_, err = mongo.coll.UpdateOne(ctx, filter, updater)
	return
}
