package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/AlimKenzza/authorization/foodService"
	"gitlab.com/AlimKenzza/authorization/pkg/dataFood"
	"strconv"
)

var jsonContentType = "application/json; charset=utf-8"
var foodRepository foodService.IFoodsRepository

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	food := foodRepository.GetFoodById(id)
	c.JSON(200, food)
}

func GetAllFoods(c *gin.Context) {
	foods := foodRepository.GetAllFoods()
	c.JSON(200, foods)
}

func DeleteFood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	food := foodRepository.GetFoodById(id)
	if food == nil {
		c.Data(400, jsonContentType, []byte("No such food with id"))
		return
	}
	if foodRepository.DeleteFood(*food) {
		c.Data(200, jsonContentType, []byte("Food deleted successfully"))
		return
	}
	c.Data(500, jsonContentType, []byte("Failed to delete food"))
}

func CreateFood(c *gin.Context) {
	food := &dataFood.Food{}
	err := c.BindJSON(food)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	if foodRepository.CreateFood(*food) {
		c.Data(200, jsonContentType, []byte("Created food \n"))
	}
	c.Data(500, jsonContentType, []byte("Failed to create food"))
}
