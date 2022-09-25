package user

import (
	"context"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user User) (err error)
	FindUserById(ctx context.Context, id string) (user *User, err error)
	FindUserByEmail(ctx context.Context, email string) (user *User, err error)
	UpdateUser(ctx context.Context, user User) (err error)
	DeleteUser(ctx context.Context, id string) (err error)
}
