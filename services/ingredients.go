package services

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/minhlong149/what-the-food/models"
)

type IngredientService struct {
	Db *sql.DB
}

func (s *IngredientService) GetAllIngredient() ([]models.Ingredient, error) {
	rows, err := s.Db.Query(`SELECT * FROM ingredients`)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var ingredients []models.Ingredient

	for rows.Next() {
		var ingredient models.Ingredient

		err = rows.Scan(&ingredient.Id, &ingredient.Name, &ingredient.Category, &ingredient.Image)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		ingredients = append(ingredients, ingredient)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return ingredients, nil
}
