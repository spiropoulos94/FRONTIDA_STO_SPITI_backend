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
		userGroup.GET("/services", controllers.UserServices)
		userGroup.GET("/:id", controllers.FindUser)
		// userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	userGroup.POST("/admin-create", CheckIfUserIsAdmin(), controllers.AdminCreateUser) // middleware to check if user is admin

	reportGroup := router.Group("/reports", CheckHeaderForJWT())
	{
		reportGroup.GET("/all", CheckIfUserIsAdmin(), controllers.ListAllReports)
		// reportGroup.GET("/", controllers.ListAvailableReports)
	}

	router.POST("/complete-signup", controllers.CompleteSignUp)
	router.POST("/login", controllers.Login)

	router.Run()
}
