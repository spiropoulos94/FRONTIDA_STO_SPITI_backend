package controllers

import (
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"
	"strconv"

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

func ListUserReports(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	userReports, err := models.GetUserReports(id)
	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"ok":      true,
		"message": "Reports Retrieved succesfully",
		"reports": userReports,
	})
}

func CreateReport(c *gin.Context) {

	// create Report from models.CreateReport
	// newReportID, err := models.CreateReport(UserID,PatientID, ReportContent, ArrivalTime, DepartureTime, AbscenceStatus)

	// if err != nil {
	// 	ErrorJSON(c, err.Error())
	// 	return
	// }

	c.JSON(200, gin.H{
		"ok": true,
		// "":         newUserID,
		"message": "Report Created",
	})
}
