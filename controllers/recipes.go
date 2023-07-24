package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/minhlong149/what-the-food/services"
)

func CheckRecipe(w http.ResponseWriter, r *http.Request) {
	dish := r.URL.Query().Get("dish")
	ingredient := r.URL.Query().Get("ingredient")

	if dish == "" || ingredient == "" {
		http.Error(w, "Missing dish or ingredient", http.StatusBadRequest)
		return
	}

	recipe, err := services.CheckRecipe(dish, ingredient)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(recipe)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
