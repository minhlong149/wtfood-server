package main

import (
	"net/http"

	"github.com/minhlong149/what-the-food/controllers"
)

func main() {
	http.HandleFunc("/api/ingredients", controllers.GetAllIngredients)
	http.HandleFunc("/api/dishes", controllers.GetRandomDish)
	http.HandleFunc("/api/recipes", controllers.CheckRecipe)

	http.ListenAndServe(":8080", nil)
}
