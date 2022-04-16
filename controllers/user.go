package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	fmt.Println("getting users!")
}
