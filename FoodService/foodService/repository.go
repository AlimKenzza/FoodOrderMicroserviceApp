package foodService

import (
	"gitlab.com/AlimKenzza/authorization/pkg/dataFood"
)

type IFoodsRepository interface {
	CreateFood(food dataFood.Food) bool
	GetAllFoods() []*dataFood.Food
	GetFoodById(id int) *dataFood.Food
	DeleteFood(order dataFood.Food) bool
}
