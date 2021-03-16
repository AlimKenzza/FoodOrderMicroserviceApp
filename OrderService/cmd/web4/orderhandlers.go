package web4

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/AlimKenzza/authorization/interfaces"
	"gitlab.com/AlimKenzza/authorization/pkg/dataOrders"
	"strconv"
)

var jsonContentType = "application/json; charset=utf-8"
var orderRepository interfaces.IOrderRepository

func GetOrderById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	order := orderRepository.GetOrderById(id)
	c.JSON(200, order)
}
func GetAllOrders(c *gin.Context) {
	orders := orderRepository.GetAllOrders()
	c.JSON(200, orders)
}

func DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	order := orderRepository.GetOrderById(id)
	if order == nil {
		c.Data(400, jsonContentType, []byte("No such user with id"))
		return
	}
	if orderRepository.DeleteOrder(*order) {
		c.Data(200, jsonContentType, []byte("Order deleted successfully"))
		return
	}
	c.Data(500, jsonContentType, []byte("Failed to delete order"))
}
func CreateOrder(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	order := orderRepository.GetOrderById(id)
	if order != nil {
		c.Data(400, jsonContentType, []byte("Order with same id exists"))
		return
	}
	food_id, err := strconv.Atoi(c.Param("food_id"))
	if err != nil || food_id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil || user_id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil || count < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect count format"))
		return
	}
	order = &dataOrders.Order{
		ID:     id,
		UserID: user_id,
		FoodID: food_id,
		Count:  count,
	}

}