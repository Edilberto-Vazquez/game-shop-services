package usecases

import (
	"context"
	"log"

	"github.com/Edilberto-Vazquez/game-shop-services/src/user/config"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/domains"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/drivers/database"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/models"
	"github.com/google/uuid"
)

type SessionService struct {
	users domains.UserRepository
}

type SessionConfig func(ss *SessionService) error

func NewSessionService(cfgs ...SessionConfig) *SessionService {
	ss := &SessionService{}
	for _, cfg := range cfgs {
		err := cfg(ss)
		if err != nil {
			log.Fatal(err)
		}
	}
	return ss
}

func WithMongoUserRepository() SessionConfig {
	return func(ss *SessionService) error {
		cr, err := database.NewMongoRepository(config.NewMongoConfig())
		if err != nil {
			return err
		}
		ss.users = cr
		return nil
	}
}

func (ss *SessionService) SignUp(ctx context.Context, user *models.Person) (uuid.UUID, error) {
	u, err := domains.NewUser(user)
	if err != nil {
		return uuid.Nil, err
	}
	ss.users.InsertUser(ctx, u)
	return u.GetID(), nil
}

func (ss *SessionService) Login(ctx context.Context, userName string) (string, error) {
	ss.users.FindUser(ctx, userName)
	return userName, nil
}
