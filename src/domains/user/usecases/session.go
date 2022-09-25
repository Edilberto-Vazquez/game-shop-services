package usecases

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Edilberto-Vazquez/game-shop-services/src/domains/shared/valueobjects"
	"github.com/Edilberto-Vazquez/game-shop-services/src/domains/user"
	"github.com/Edilberto-Vazquez/game-shop-services/src/domains/user/drivers/mongodb"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

var (
	validate *validator.Validate
)

type UserSession struct {
	users     user.UserRepository
	JWTSecret string
}

type SessionConfig func(ss *UserSession) error

func NewSession(cfgs ...SessionConfig) *UserSession {
	ss := &UserSession{}
	for _, cfg := range cfgs {
		err := cfg(ss)
		if err != nil {
			log.Fatal(err)
		}
	}
	return ss
}

func WithMongoUserRepository() SessionConfig {
	return func(ss *UserSession) error {
		repo, err := mongodb.NewMongoDBRepository()
		if err != nil {
			return err
		}
		ss.users = repo
		ss.JWTSecret = os.Getenv("DEV_JWT_SECRET")
		return nil
	}
}

func (ss *UserSession) SignUp(ctx context.Context, u user.User) (statusCode int, err error) {
	validate = validator.New()
	err = validate.Struct(u)
	if err != nil {
		return http.StatusBadRequest, err
	}
	emailExist, err := ss.users.FindUserByEmail(ctx, u.Email)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return http.StatusInternalServerError, errors.New("something happened while registering the user")
	}
	if emailExist != nil {
		return http.StatusInternalServerError, errors.New("this email has been registered")
	}
	hashedPwd, err := u.HashPassword()
	if err != nil {
		return http.StatusInternalServerError, errors.New("something happened while registering the user")
	}
	u.Password = hashedPwd
	err = ss.users.InsertUser(ctx, u)
	if err != nil {
		return http.StatusInternalServerError, errors.New("something happened while registering the user")
	}
	return http.StatusOK, nil
}

func (ss *UserSession) Login(ctx context.Context, email, password string) (tokenString string, statusCode int, err error) {
	u, err := ss.users.FindUserByEmail(ctx, email)
	if err != nil {
		return "", http.StatusBadRequest, errors.New("this email does not exist")
	}
	err = u.PasswordMatch(password)
	if err != nil {
		return "", http.StatusBadRequest, errors.New("the password is wrong")
	}
	claims := valueobjects.AppClaims{
		UserEmail: u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(ss.JWTSecret))
	if err != nil {
		return "", http.StatusInternalServerError, errors.New("could not login please try or check your information")
	}
	return tokenString, http.StatusOK, nil
}
