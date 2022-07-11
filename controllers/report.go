package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	report := models.Report{}
	json.Unmarshal(jsonData, &report)

	fmt.Printf("%+v\n", report)

	// check if patient already exists
	patient, err := models.GetPatientByAMKA(report.Patient.Patient_AMKA)

	if err != nil {
		if err.Error() == "404" {
			fmt.Println("Patient does not exist, will create one")
			createdPatientID, _ := models.SavePatient(report.Patient)
			fmt.Println("CREATED PATIENT ID =>", createdPatientID)
		}
	}

	fmt.Println("patient => ", patient)

	// // create Report from models.SaveReport
	// newReportID, err := models.SaveReport(report.User_id, report.Patient_id, report.ReportContent, report.ArrivalTime, report.DepartureTime, report.AbscenceStatus)

	// if err != nil {
	// 	ErrorJSON(c, err.Error())
	// 	return
	// }

	c.JSON(200, gin.H{
		"ok": true,
		// "newReportID": newReportID,
		"message": "Report Created",
	})
}
