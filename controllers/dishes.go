package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/minhlong149/what-the-food/services"
)

func GetRandomDish(w http.ResponseWriter, r *http.Request) {
	dish, err := services.GetRandomDish()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(dish)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
