package controllers

import (
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"

	"github.com/gin-gonic/gin"
)

func ListAllReports(c *gin.Context) {
	var reports []models.Report

	reports, err := models.GetAllReports()

	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "Reports Retrieved succesfully",
		"reports": reports,
	})
}

// func ListAvailableReports(c *gin.Context) {
// 	// list all reports that the signed user can read
// }
