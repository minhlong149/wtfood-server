package services

import (
	"database/sql"
	"log"

	"github.com/minhlong149/what-the-food/config"
)

type Recipe struct {
	Dish       *Dish       `json:"dish,omitempty"`
	Ingredient *Ingredient `json:"ingredient,omitempty"`
	Correct    bool        `json:"correct"`
}

func CheckRecipe(dishId string, ingredientId string) (recipe Recipe, err error) {
	query := `
		SELECT
			d.id, d.name, d.category, d.image,
			i.id, i.name, i.category, i.image
		FROM recipes r
		JOIN dishes d ON d.id = r.dish_id
		JOIN ingredients i ON i.id = r.ingredient_id
		WHERE dish_id = $1 AND ingredient_id = $2
	`
	row := config.Db.QueryRow(query, dishId, ingredientId)

	dish := Dish{}
	ingredient := Ingredient{}

	err = row.Scan(
		&dish.Id, &dish.Name, &dish.Category, &dish.Image,
		&ingredient.Id, &ingredient.Name, &ingredient.Category, &ingredient.Image,
	)

	switch err {
	case sql.ErrNoRows:
		recipe.Correct = false
		err = nil

	case nil:
		recipe.Dish = &dish
		recipe.Ingredient = &ingredient
		recipe.Correct = true

	default:
		log.Println(err)
		return
	}

	return
}
