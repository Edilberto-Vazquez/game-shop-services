package services

import (
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/usecases"
)

type Services struct {
	SessionService *usecases.Session
}

func NewServices() *Services {
	return &Services{
		SessionService: usecases.NewSession(usecases.WithMongoUserRepository()),
	}
}
