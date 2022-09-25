package user

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string    `json:"email" binding:"required" validate:"required" bson:"email"`
	UserName  string    `json:"userName" binding:"required" validate:"required" bson:"user_name"`
	CountryId string    `json:"countryId" binding:"required" validate:"required" bson:"country_id"`
	Password  string    `json:"password" binding:"required" validate:"required" bson:"password"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deletedAt" bson:"deleted_at,omitempty"`
}

func NewUser(email, userName, countryId, password string) *User {
	return &User{
		Email:     email,
		UserName:  userName,
		CountryId: countryId,
		Password:  password,
	}
}

func (u *User) HashPassword() (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(fmt.Errorf("[USER] HashPassword: %w", err))
		return "", err
	}
	return string(hashedPwd), err
}

func (u *User) PasswordMatch(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Println(fmt.Errorf("[USER] PasswordMatch: %w", err))
		return err
	}
	return err
}
