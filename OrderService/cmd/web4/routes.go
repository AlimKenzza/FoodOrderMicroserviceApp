package web4

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/orders/:id", GetOrderById)
	router.GET("/orders", GetAllOrders)
	router.DELETE("/orders/:id", DeleteOrder)
	router.POST("/orders/:id", CreateOrder)
	return router
}