package usecases

import (
	"context"
	"net/http"

	"github.com/Edilberto-Vazquez/game-shop-services/session/models"
)

func SignUp(insertUser func(ctx context.Context, user *models.User) (err error)) func(ctx context.Context, user *models.User) (int, error) {
	return func(ctx context.Context, user *models.User) (int, error) {
		if err := insertUser(ctx, user); err != nil {
			return http.StatusInternalServerError, err
		} else {
			return http.StatusOK, nil
		}
	}
}
