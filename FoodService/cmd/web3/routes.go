package main

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/foods/:id", Get)
	router.GET("/foods", GetAllFoods)
	//router.PUT("/users/:id", Update)
	router.DELETE("/foods/:id", DeleteFood)
	router.POST("/foods", CreateFood)
	return router
}
