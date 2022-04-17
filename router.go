package main

import (
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {

	router := gin.Default()

	userGroup := router.Group("/user")
	{
		// user group handlers
		userGroup.GET("/", controllers.ListUsers)
		// userGroup.GET("/:id", controllers.FindUser)
		userGroup.POST("/", controllers.AdminCreateUser)
		// userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	router.POST("/signup", controllers.SignUp)

	router.Run()
}
