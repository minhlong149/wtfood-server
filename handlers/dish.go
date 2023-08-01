package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/minhlong149/what-the-food/models"
)

type DishHandler struct {
	DishService models.DishService
}

func (h *DishHandler) GetRandomDish(w http.ResponseWriter, r *http.Request) {
	dish, err := h.DishService.GetRandomDish()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(dish)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
