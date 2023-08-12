package service

func (r *Repository) GetAllIngredients() (ingredients []Ingredient, err error) {
	err = r.Db.Find(&ingredients).Error
	return ingredients, err
}
