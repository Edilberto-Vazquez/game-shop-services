package user

import (
	"context"

	"github.com/Edilberto-Vazquez/game-shop-services/src/user/domains"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/models"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/usecases"
)

type UserService struct {
	UserModel   models.User
	LoginModel  models.Login
	SignUpModel models.SignUp
	SignUp      func(ctx context.Context, signup *models.SignUp) (userId string, err error)
	Login       func(ctx context.Context, userId string) (user *models.User, err error)
}

func NewUserService() *UserService {
	domains.NewUserDomain()
	return &UserService{
		SignUp: usecases.SignUp,
		Login:  usecases.Login,
	}
}
