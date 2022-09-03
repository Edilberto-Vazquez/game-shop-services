package models

type User struct {
	Id        string `json:"id"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	CountryId string `json:"countryId"`
	Salt      string `json:"salt"`
	Hash      string `json:"hash"`
}

type HandleUserSession interface {
}
