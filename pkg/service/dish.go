package service

import (
	"gorm.io/gorm"

	"wtfood/pkg/model"
)

func (r *Repository) GetRandomDish() (dish Dish, err error) {
	err = r.Db.Model(&Dish{}).Order("RANDOM()").First(&dish).Error
	if err != nil {
		return Dish{}, err
	}

	err = r.Db.Model(&dish).Association("Ingredients").Find(&dish.Ingredients)
	if err != nil {
		return Dish{}, err
	}

	return dish, nil
}

func (r *Repository) GetAllDishes() (dishes []Dish, err error) {
	err = r.Db.Model(&Dish{}).Find(&dishes).Error
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (r *Repository) GetDishById(dishId string) (dish Dish, err error) {
	err = r.Db.Model(&Dish{}).First(&dish, dishId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Dish{}, model.ErrDishNotFound
		}
		return Dish{}, err
	}

	err = r.Db.Model(&dish).Association("Ingredients").Find(&dish.Ingredients)
	if err != nil {
		return Dish{}, err
	}

	return dish, nil
}
