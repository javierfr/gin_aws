package routes

import (
	"gin_aws/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)
		api.POST("/users", controllers.CreateUser)
		api.GET("/users/:id", controllers.GetUser)
		api.DELETE("/users/:id", controllers.DeleteUser)
	}

	return router
}
