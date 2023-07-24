package services

import (
	"log"

	"github.com/minhlong149/what-the-food/config"
)

type Dish struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Image    string `json:"image"`
}

func GetRandomDish() (dish Dish, err error)  {
	row := config.Db.QueryRow(`SELECT * FROM dishes ORDER BY random() LIMIT 1`)
	err = row.Scan(&dish.Id, &dish.Name, &dish.Category, &dish.Image)

	if err != nil {
		log.Println(err)
		return
	}

	return
}
