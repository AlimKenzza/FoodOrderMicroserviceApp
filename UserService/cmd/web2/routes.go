package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users/:id", Get)
	router.GET("/users", GetAllUsers)
	//router.PUT("/users/:id", Update)
	router.DELETE("/users/:id", DeleteUser)
	return router
}
