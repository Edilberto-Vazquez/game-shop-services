package user

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

var (
	validate *validator.Validate
)

type Session struct {
	users     UserRepository
	JWTSecret string
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
		conf, err := NewConfig()
		if err != nil {
			return err
		}
		repo, err := NewMongoDBRepository(conf)
		if err != nil {
			return err
		}
		ss.users = repo
		ss.JWTSecret = os.Getenv("DEV_JWT_SECRET")
		return nil
	}
}

func (ss *Session) SignUp(ctx context.Context, user *Person) (id string, err error) {
	validate = validator.New()
	err = validate.Struct(user)
	if err != nil {
		return "", errors.New("a user has to have a valid person")
	}
	hashedPwd, err := HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPwd
	id, err = ss.users.InsertUser(ctx, NewUser(user))
	if err != nil {
		return "", err
	}
	return id, nil
}

func (ss *Session) Login(ctx context.Context, login Login) (LoginResponse, error) {
	user, err := ss.users.FindUserByEmail(ctx, login.Email)
	if err != nil {
		return LoginResponse{}, err
	}
	err = PasswordMatch(user.GetPassword(), login.Password)
	if err != nil {
		return LoginResponse{}, err
	}
	claims := AppClaims{
		UserEmail: user.GetEmail(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(ss.JWTSecret))
	if err != nil {
		return LoginResponse{}, err
	}
	return LoginResponse{
		Email: user.GetEmail(),
		Token: tokenString,
	}, nil
}
