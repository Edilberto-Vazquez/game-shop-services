package domains

import (
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/models"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

type User struct {
	person *models.Person
}

func NewUser(user *models.Person) User {
	return User{
		person: user,
	}
}

func (u *User) GetID() string {
	return u.person.ID
}

func (u *User) SetID(id string) {
	u.person.ID = id
}

func (u *User) GetUserName() string {
	return u.person.UserName
}

func (u *User) SetUserName(userName string) {
	u.person.UserName = userName
}

func (u *User) GetEmail() string {
	return u.person.Email
}

func (u *User) SetEmail(email string) {
	u.person.Email = email
}

func (u *User) GetCountryId() string {
	return u.person.CountryId
}

func (u *User) SetCountryId(countryId string) {
	u.person.CountryId = countryId
}

func (u *User) GetPassword() string {
	return u.person.Password
}

func (u *User) SetPassword(password string) {
	u.person.Password = password
}
