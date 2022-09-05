package domains

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user User) (userId string, err error)
	FindUser(ctx context.Context, userId string) (user User, err error)
	UpdateUser(ctx context.Context, user User) (err error)
	DeleteUser(ctx context.Context, userId uuid.UUID) (err error)
}
