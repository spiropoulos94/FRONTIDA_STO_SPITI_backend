package main

import (
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {

	router := gin.Default()

	userGroup := router.Group("/user", CheckHeaderForJWT())
	{
		// user group handlers
		userGroup.GET("/", controllers.ListUsers)
		// userGroup.GET("/:id", controllers.FindUser)
		// userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	userGroup.POST("/admin-create", CheckIfUserIsAdmin(), controllers.AdminCreateUser) // middleware to check if user is admin

	router.POST("/complete-signup", controllers.CompleteSignUp)
	router.POST("/login", controllers.Login)

	router.Run()
}
