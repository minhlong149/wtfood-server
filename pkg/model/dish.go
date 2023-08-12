package model

type Dish struct {
	Id          string       `json:"id" gorm:"primary_key"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"-" gorm:"many2many:recipes;"`
}
