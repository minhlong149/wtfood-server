package model

import (
	"errors"
)

var (
	ErrDishNotFound = errors.New("i couldn't find that dish, maybe you should try another one?")
	ErrMissingDishId = errors.New("i can't check a recipe without knowing which dish you want to cook")
)

var (
	ErrIngredientNotFound = errors.New("i couldn't find that ingredient, maybe you should try another one?")
	ErrMissingIngredientId = errors.New("i can't check a recipe without knowing which ingredient you want to use")
)

var (
	ErrInvalidCredentials = errors.New("i can't let you in without knowing who you are, please enter your username and password")
	ErrWrongCredentials   = errors.New("that's not the password I was expecting, are you sure you're not trying to log in to someone else's account?")
)

var (
	ErrInvalidAuthHeader = errors.New("your authorization header is so invalid, it makes me question your entire life")
	ErrInvalidToken      = errors.New("your token is so invalid, it got rejected from Hogwarts")
	ErrUnauthorized      = errors.New("looks like you're not authorized to access this site, maybe you should try logging in as someone else?")
)
