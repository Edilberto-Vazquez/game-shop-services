package services

import "github.com/Edilberto-Vazquez/game-shop-services/src/user"

type Services struct {
	SessionService *user.Session
}

func NewServices() *Services {
	return &Services{
		SessionService: user.NewSession(user.WithMongoUserRepository()),
	}
}
