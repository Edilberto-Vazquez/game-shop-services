package usecases

import (
	"context"
	"log"

	"github.com/Edilberto-Vazquez/game-shop-services/src/user/models"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/repository"
)

func SignUp(ctx context.Context, signup *models.SignUp) (userId string, err error) {
	user := models.User{
		UserName:  signup.UserName,
		Email:     signup.Email,
		CountryId: signup.CountryId,
		Salt:      signup.Salt,
		Hash:      signup.Hash,
	}
	if userId, err = repository.InsertUser(ctx, &user); err != nil {
		return "", err
	} else {
		log.Println(userId)
		return userId, nil
	}
}

func Login(ctx context.Context, userId string) (user *models.User, err error) {
	if user, err = repository.FindUser(ctx, userId); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
