package models

type Ingredient struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Image    string `json:"image"`
}

type IngredientService interface {
	GetAllIngredient() ([]Ingredient, error)
}
