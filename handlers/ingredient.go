package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/minhlong149/what-the-food/models"
)

type IngredientHandler struct {
	IngredientService models.IngredientService
}

func (h *IngredientHandler) GetAllIngredient(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.IngredientService.GetAllIngredient()
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
