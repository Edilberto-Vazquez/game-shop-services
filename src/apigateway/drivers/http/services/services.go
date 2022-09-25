package services

import "github.com/Edilberto-Vazquez/game-shop-services/src/domains/user/usecases"

type Services struct {
	UserSessionService *usecases.UserSession
}

func NewServices() *Services {
	return &Services{
		UserSessionService: usecases.NewSession(usecases.WithMongoUserRepository()),
	}
}
