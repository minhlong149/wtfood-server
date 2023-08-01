package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/minhlong149/what-the-food/models"
)

type RecipeHandler struct {
	RecipeService models.RecipeService
}

func (h *RecipeHandler) CheckRecipe(w http.ResponseWriter, r *http.Request) {
	dish := r.URL.Query().Get("dish")
	ingredient := r.URL.Query().Get("ingredient")

	if dish == "" || ingredient == "" {
		http.Error(w, "Missing dish or ingredient", http.StatusBadRequest)
		return
	}

	ingredients, err := h.RecipeService.CheckRecipe(dish, ingredient)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(ingredients)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
