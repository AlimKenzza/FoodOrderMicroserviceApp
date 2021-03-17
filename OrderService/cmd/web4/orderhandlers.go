package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/AlimKenzza/authorization/interfaces"
	"gitlab.com/AlimKenzza/authorization/pkg/auth"
	"gitlab.com/AlimKenzza/authorization/pkg/dataOrders"
	"strconv"
)

var jsonContentType = "application/json; charset=utf-8"
var orderRepository interfaces.IOrderRepository

func GetOrderById(c *gin.Context) {
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
func CreateOrder(c *gin.Context) {
	err := auth.TokenValid(c.Request)
	if err != nil {
		c.Data(403, jsonContentType, []byte("Invalid token"))
		return
	}
	userId, err := auth.ExtractTokenMetadata(c.Request)
	if userId == 0 {
		c.Data(401, jsonContentType, []byte("Unauthorized. Please, sign in"))
		return
	}
	order := &dataOrders.Order{}
	err = c.BindJSON(order)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	order.UserID = userId
	if orderRepository.CreateOrder(*order) {
		c.Data(200, jsonContentType, []byte("Created order \n"))
	}
	c.Data(500, jsonContentType, []byte("Failed to create order"))

}
