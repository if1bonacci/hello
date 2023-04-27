package models

type UserType struct {
	Id       string
	UserName string `json:"userName"`
	Password string `json:"password"`
}

var User = new(UserType)
