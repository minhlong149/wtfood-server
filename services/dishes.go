package services

import (
	"log"

	"github.com/minhlong149/what-the-food/config"
)

type Dish struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	Ingredients int    `json:"ingredients"`
}

func GetRandomDish() (dish Dish, err error) {
	row := config.Db.QueryRow(`
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

	err = row.Scan(&dish.Id, &dish.Name, &dish.Category, &dish.Image, &dish.Ingredients)

	if err != nil {
		log.Println(err)
		return
	}

	return
}
