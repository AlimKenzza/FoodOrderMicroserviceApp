package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/AlimKenzza/authorization/cmd/web/handlers"
	"gitlab.com/AlimKenzza/authorization/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// Middlewares
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORSMiddleware())
	router.POST("/register", handlers.Create)
	router.POST("/login", handlers.Login)
	return router
}
