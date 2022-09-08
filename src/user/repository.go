package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user User) (err error)
	FindUserById(ctx context.Context, id primitive.ObjectID) (user User, err error)
	FindUserByEmail(ctx context.Context, email string) (user User, err error)
	UpdateUser(ctx context.Context, user User) (err error)
	DeleteUser(ctx context.Context, id primitive.ObjectID) (err error)
}
