package models

import (
	"errors"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UserService interface {
	AccountIsValid(Account) (error)
	Register(Account) (User, error)
}

var ErrUsernameExists = errors.New("username exists")

var ErrInvalidUsername = errors.New("invalid username")

var ErrInvalidPassword = errors.New("invalid password")
