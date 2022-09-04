package repository

import (
	"context"

	"github.com/Edilberto-Vazquez/game-shop-services/src/user/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) (userId string, err error)
	FindUser(ctx context.Context, userId string) (user *models.User, err error)
	UpdateUser(ctx context.Context, user *models.User) (err error)
	DeleteUser(ctx context.Context, userId string) (err error)
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) (userId string, err error) {
	return implementation.InsertUser(ctx, user)
}

func FindUser(ctx context.Context, userId string) (user *models.User, err error) {
	return implementation.FindUser(ctx, userId)
}

func UpdateUser(ctx context.Context, user *models.User) (err error) {
	return implementation.UpdateUser(ctx, user)
}

func DeleteUser(ctx context.Context, userId string) (err error) {
	return implementation.DeleteUser(ctx, userId)
}
