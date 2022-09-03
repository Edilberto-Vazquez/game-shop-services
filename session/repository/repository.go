package repository

import (
	"context"

	"github.com/Edilberto-Vazquez/game-shop-services/session/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) (err error) {
	return implementation.InsertUser(ctx, user)
}
