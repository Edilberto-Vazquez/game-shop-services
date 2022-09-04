package models

type Login struct {
	UserName string `json:"userName"`
}

type SignUp struct {
	UserName  string `json:"userName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	CountryId string `json:"countryId" binding:"required"`
	Salt      string `json:"salt" binding:"required"`
	Hash      string `json:"hash" binding:"required"`
}

type User struct {
	Id        string `json:"id"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	CountryId string `json:"countryId"`
	Salt      string `json:"salt"`
	Hash      string `json:"hash"`
}
