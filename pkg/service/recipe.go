package service

func (r *Repository) DishHasIngredient(dishId, ingredientId string) (dish Dish, ingredient Ingredient, exist bool, err error) {
	ingredient, err = r.GetIngredientById(ingredientId)
	if err != nil {
		return Dish{}, Ingredient{}, false, err
	}

	dish, err = r.GetDishById(dishId)
	if err != nil {
		return Dish{}, Ingredient{}, false, err
	}

	for _, dishIngredient := range dish.Ingredients {
		if dishIngredient.Id == ingredient.Id {
			return dish, ingredient, true, nil
		}
	}

	return dish, ingredient, false, nil
}
