package service

import (
	"gorm.io/gorm"

	"wtfood/pkg/model"
)

type Repository struct {
	Db *gorm.DB
}

type Ingredient = model.Ingredient

type Dish = model.Dish

type Account = model.Account

type Service interface {
	GetAllIngredients() ([]Ingredient, error)
	GetAllDishes() ([]Dish, error)
	GetDishById(string) (Dish, error)
	GetRandomDish() (Dish, error)
	Login(Account, string) (string, error)
}
