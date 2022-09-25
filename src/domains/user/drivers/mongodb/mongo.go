package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Edilberto-Vazquez/game-shop-services/src/config"
	"github.com/Edilberto-Vazquez/game-shop-services/src/domains/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	gameShop *mongo.Database
	users    *mongo.Collection
}

func NewMongoDBRepository() (*MongoRepository, error) {
	mongoUri := os.Getenv(config.Env + "MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Println(fmt.Errorf("[MONGO_DRIVER] NewMongoDBRepository: %W", err))
		return nil, err
	}
	clientDB := client.Database("game-shop")
	clientColl := clientDB.Collection("users")
	return &MongoRepository{gameShop: clientDB, users: clientColl}, nil
}

func (m *MongoRepository) InsertUser(ctx context.Context, user user.User) (err error) {
	user.CreatedAt = time.Now()
	_, err = m.users.InsertOne(ctx, user)
	if err != nil {
		log.Println(fmt.Errorf("[MONGO_DRIVER] FindUserById: %w", err))
		return err
	}
	return nil
}

func (m *MongoRepository) FindUserById(ctx context.Context, id string) (user *user.User, err error) {
	filter := bson.M{"_id": id, "deleted_at": bson.M{"$exist": false}}
	err = m.users.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Println(fmt.Errorf("[MONGO_DRIVER] FindUserById: %w", err))
		return nil, err
	}
	return user, nil
}

func (m *MongoRepository) FindUserByEmail(ctx context.Context, email string) (user *user.User, err error) {
	filter := bson.M{"email": email, "deleted_at": bson.M{"$exists": false}}
	err = m.users.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Println(fmt.Errorf("[MONGO_DRIVER] FindUserByEmail: %w", err))
		return nil, err
	}
	return user, nil
}

func (m *MongoRepository) UpdateUser(ctx context.Context, user user.User) (err error) {
	filter := bson.M{"_id": user.ID, "deleted_at": bson.M{"$exist": false}}
	user.UpdatedAt = time.Now()
	_, err = m.users.UpdateOne(ctx, filter, user)
	if err != nil {
		log.Println(fmt.Errorf("[MONGO_DRIVER] UpdateUser: %w", err))
		return err
	}
	return nil
}

func (mongo *MongoRepository) DeleteUser(ctx context.Context, id string) (err error) {
	filter := bson.M{"_id": id, "deleted_at": bson.M{"$exist": false}}
	updater := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "deleted_at", Value: time.Now()},
				{Key: "updated_at", Value: time.Now()},
			}},
	}
	_, err = mongo.users.UpdateOne(ctx, filter, updater)
	if err != nil {
		log.Println(fmt.Errorf("[MONGO_DRIVER] UpdateUser: %w", err))
		return err
	}
	return nil
}
