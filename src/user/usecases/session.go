package usecases

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/Edilberto-Vazquez/game-shop-services/src/user/config"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/domains"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/drivers/database"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/models"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

var (
	validate *validator.Validate
)

type Session struct {
	users domains.UserRepository
}

type loginResponse struct {
	userName string `json:userName`
	token    string `json:"token"`
}

type SessionConfig func(ss *Session) error

func NewSession(cfgs ...SessionConfig) *Session {
	ss := &Session{}
	for _, cfg := range cfgs {
		err := cfg(ss)
		if err != nil {
			log.Fatal(err)
		}
	}
	return ss
}

func WithMongoUserRepository() SessionConfig {
	return func(ss *Session) error {
		repo, err := database.NewMongoDBRepository(config.NewMongoConfig())
		if err != nil {
			return err
		}
		ss.users = repo
		return nil
	}
}

func (ss *Session) SignUp(ctx context.Context, user *models.Person) (id string, err error) {
	validate = validator.New()
	err = validate.Struct(user)
	if err != nil {
		return "", errors.New("a user has to have a valid person")
	}
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPwd
	id, err = ss.users.InsertUser(ctx, domains.NewUser(user))
	if err != nil {
		return "", err
	}
	return id, nil
}

func (ss *Session) Login(ctx context.Context, email, password string) (loginResponse, error) {
	user, err := ss.users.FindUserByEmail(ctx, email)
	if err != nil {
		return loginResponse{}, err
	}
	err = utils.PasswordMatch(user.GetPassword(), password)
	if err != nil {
		return loginResponse{}, err
	}
	claims := models.AppClaims{
		UserID: user.GetID(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour * 24)),
		},
	}
	signinKey := os.Getenv("DEV_JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signinKey))
	if err != nil {
		return loginResponse{}, err
	}
	return loginResponse{
		userName: user.GetUserName(),
		token:    tokenString,
	}, nil
}
