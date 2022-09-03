package session

import (
	"context"

	"github.com/Edilberto-Vazquez/game-shop-services/session/config"
	"github.com/Edilberto-Vazquez/game-shop-services/session/domains"
	"github.com/Edilberto-Vazquez/game-shop-services/session/drivers"
	"github.com/Edilberto-Vazquez/game-shop-services/session/models"
	"github.com/Edilberto-Vazquez/game-shop-services/session/usecases"
)

type SessionService struct {
	SignUp func(ctx context.Context, user *models.User) (int, error)
}

func NewSessionService() *SessionService {
	var db *drivers.MongoDB = drivers.NewMongoDB(config.Config())
	return &SessionService{
		SignUp: usecases.SignUp(domains.InsertUser(db)),
	}
}
