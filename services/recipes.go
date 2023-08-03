package services

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/minhlong149/what-the-food/models"
)

type RecipeService struct {
	Db *sql.DB
}

func (s *RecipeService) CheckRecipe(dishId string, ingredientId string) (models.Recipe, error) {
	query := `
		SELECT d.id, d.name, i.id, i.name, i.image
		FROM recipes r
		JOIN dishes d ON d.id = r.dish_id
		JOIN ingredients i ON i.id = r.ingredient_id
		WHERE dish_id = $1 AND ingredient_id = $2
	`
	dish := models.Dish{}
	ingredient := models.Ingredient{}
	recipe := models.Recipe{}

	row := s.Db.QueryRow(query, dishId, ingredientId)

	err := row.Scan(
		&dish.Id, &dish.Name,
		&ingredient.Id, &ingredient.Name, &ingredient.Image,
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
		return recipe, err
	}

	return recipe, nil
}
