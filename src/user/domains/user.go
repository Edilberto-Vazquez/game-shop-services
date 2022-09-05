package domains

import (
	"errors"

	"github.com/Edilberto-Vazquez/game-shop-services/src/user/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	validate *validator.Validate
)

type User struct {
	person *models.Person
}

func NewUser(user *models.Person) (User, error) {
	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		return User{}, errors.New("a user has to have a valid person")
	}
	return User{
		person: user,
	}, nil
}

func (u *User) GetID() uuid.UUID {
	return u.person.ID
}

func (u *User) SetID(id uuid.UUID) {
	if u.person == nil {
		u.person = &models.Person{}
	}
	u.person.ID = id
}

func (u *User) GetUserName() string {
	return u.person.UserName
}

func (u *User) SetUserName(userName string) {
	if u.person == nil {
		u.person = &models.Person{}
	}
	u.person.UserName = userName
}

func (u *User) GetEmail() string {
	return u.person.Email
}

func (u *User) SetEmail(Email string) {
	if u.person == nil {
		u.person = &models.Person{}
	}
	u.person.Email = Email
}

func (u *User) GetCountryId() string {
	return u.person.CountryId
}

func (u *User) SetCountryId(CountryId string) {
	if u.person == nil {
		u.person = &models.Person{}
	}
	u.person.CountryId = CountryId
}

func (u *User) GetSalt() string {
	return u.person.Salt
}

func (u *User) SetSalt(Salt string) {
	if u.person == nil {
		u.person = &models.Person{}
	}
	u.person.Salt = Salt
}

func (u *User) GetHash() string {
	return u.person.Hash
}

func (u *User) SetHash(Hash string) {
	if u.person == nil {
		u.person = &models.Person{}
	}
	u.person.Hash = Hash
}
