package model

type Ingredient struct {
	Id    string `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
