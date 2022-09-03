package repository

import (
	"context"

	"github.com/Edilberto-Vazquez/game-shop-services/session/models"
)

type SessionRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
}

var implementation SessionRepository

func SetRepository(repository SessionRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) (err error) {
	return implementation.InsertUser(ctx, user)
}
