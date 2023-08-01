package models

type Dish struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	Ingredients int    `json:"ingredients"`
}

type DishService interface {
	GetRandomDish() (Dish, error)
}
