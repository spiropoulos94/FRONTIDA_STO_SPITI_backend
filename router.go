package main

import (
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {

	r := gin.Default()

	userGroup := r.Group("/user")
	{
		// user group handlers
		userGroup.GET("/", controllers.GetUsers)
		// userGroup.GET("/:id", controllers.FindUser)
		// userGroup.POST("/", controllers.CreateUser)
		// userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	r.Run()
}
