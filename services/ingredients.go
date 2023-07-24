package services

import (
	"database/sql"
	"log"


	"github.com/minhlong149/what-the-food/config"
)

type Ingredient struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Image    string `json:"image"`
}

func GetAllIngredients() (ingredients []Ingredient, err error) {
	var rows *sql.Rows

	rows, err = config.Db.Query(`SELECT * FROM ingredients`)
	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var ingredient Ingredient

		err = rows.Scan(&ingredient.Id, &ingredient.Name, &ingredient.Category, &ingredient.Image)
		if err != nil {
			log.Println(err)
			return
		}

		ingredients = append(ingredients, ingredient)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return
	}

	return
}
