package services

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/minhlong149/what-the-food/models"
)

type DishService struct {
	Db *sql.DB
}

func (s *DishService) GetRandomDish() (models.Dish, error) {
	row := s.Db.QueryRow(`
		SELECT
			d.id,
			d.name,
			d.category,
			d.image,
			COUNT(ingredient_id) as ingredients
		FROM dishes d
		JOIN recipes r ON r.dish_id = d.id
		GROUP BY d.id
		ORDER BY random() LIMIT 1
	`)

	dish := models.Dish{}

	err := row.Scan(&dish.Id, &dish.Name, &dish.Category, &dish.Image, &dish.Ingredients)
	if err != nil {
		log.Println(err)
		return dish, err
	}

	return dish, nil
}
