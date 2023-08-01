package models

type Recipe struct {
	Dish       *Dish       `json:"dish,omitempty"`
	Ingredient *Ingredient `json:"ingredient,omitempty"`
	Correct    bool        `json:"correct"`
}

type RecipeService interface {
	CheckRecipe(string, string) (Recipe, error)
}
