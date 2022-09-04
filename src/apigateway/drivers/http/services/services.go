package services

import "github.com/Edilberto-Vazquez/game-shop-services/src/user"

type Services struct {
	UserService user.UserService
}

func NewServices() *Services {
	return &Services{
		UserService: *user.NewUserService(),
	}
}
