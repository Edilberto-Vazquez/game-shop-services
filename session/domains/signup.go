package domains

import (
	"context"

	"github.com/Edilberto-Vazquez/game-shop-services/session/models"
	"github.com/Edilberto-Vazquez/game-shop-services/session/repository"
)

func InsertUser(db repository.SessionRepository) func(ctx context.Context, user *models.User) (err error) {
	return func(ctx context.Context, user *models.User) error {
		return db.InsertUser(ctx, user)
	}
}
