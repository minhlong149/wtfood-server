package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/minhlong149/what-the-food/services"
)

func GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := services.GetAllIngredients()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(ingredients)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
