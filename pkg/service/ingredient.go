package service

import (
	"gorm.io/gorm"

	"wtfood/pkg/model"
)

func (r *Repository) GetAllIngredients() (ingredients []Ingredient, err error) {
	err = r.Db.Find(&ingredients).Error
	return ingredients, err
}

func (r *Repository) GetIngredientById(ingredientId string) (ingredient Ingredient, err error) {
	err = r.Db.First(&ingredient, ingredientId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Ingredient{}, model.ErrIngredientNotFound
		}
		return Ingredient{}, err
	}
	return ingredient, nil
}
