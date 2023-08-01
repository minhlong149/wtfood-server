package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/minhlong149/what-the-food/handlers"
	"github.com/minhlong149/what-the-food/services"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Fatal(err)
	}

	ingredientHandler := &handlers.IngredientHandler{
		IngredientService: &services.IngredientService{Db: db},
	}

	dishHandler := &handlers.DishHandler{
		DishService: &services.DishService{Db: db},
	}

	recipeHandler := &handlers.RecipeHandler{
		RecipeService: &services.RecipeService{Db: db},
	}

	http.HandleFunc("/api/ingredients", ingredientHandler.GetAllIngredient)
	http.HandleFunc("/api/dishes", dishHandler.GetRandomDish)
	http.HandleFunc("/api/recipes", recipeHandler.CheckRecipe)

	http.ListenAndServe(":8080", nil)
}
