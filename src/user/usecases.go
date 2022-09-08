package user

import (
	"context"
	"log"
	"net/http"
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

func (ss *Session) SignUp(ctx context.Context, user Person) SignUpResponse {
	newUser := NewUser(&user)
	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		return newUser.ToSignUpResponse(http.StatusBadRequest, err)
	}
	userWithMail, err := ss.users.FindUserByEmail(ctx, user.Email)
	if err == nil || userWithMail.GetEmail() != "" {
		return newUser.ToSignUpResponse(http.StatusBadRequest, err)
	}
	hashedPwd, err := HashPassword(user.Password)
	if err != nil {
		return newUser.ToSignUpResponse(http.StatusInternalServerError, err)
	}
	user.Password = hashedPwd
	err = ss.users.InsertUser(ctx, newUser)
	if err != nil {
		return newUser.ToSignUpResponse(http.StatusInternalServerError, err)
	}
	return newUser.ToSignUpResponse(http.StatusOK, nil)
}

func (ss *Session) Login(ctx context.Context, login Login) LoginResponse {
	user, err := ss.users.FindUserByEmail(ctx, login.Email)
	if err != nil {
		return user.ToLoginResponse("", http.StatusBadRequest, err)
	}
	err = PasswordMatch(user.GetPassword(), login.Password)
	if err != nil {
		return user.ToLoginResponse("", http.StatusBadRequest, err)
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
		return user.ToLoginResponse("", http.StatusInternalServerError, err)
	}
	return user.ToLoginResponse(tokenString, http.StatusOK, nil)
}
