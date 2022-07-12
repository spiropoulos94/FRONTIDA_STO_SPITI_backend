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

	fmt.Println("-- CREATE REPORT CONTROLLER RUN")

	// err := errors.New("test error")
	// ErrorJSON(c, err.Error())
	// return

	var createdPatientID *int64
	var createdAddressID *int64

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	report := models.Report{}
	json.Unmarshal(jsonData, &report)

	if report.Patient.Patient_AMKA == 0 || report.Patient.Patient_AMKA > 999999999 {
		ErrorJSON(c, "Use a valid AMKA")
		return
	}

	// check if patient already exists
	patient, err := models.GetPatientByAMKA(report.Patient.Patient_AMKA)

	if patient == nil || err != nil {
		createdAddressID, err = models.SaveAddress(report.Patient.Address.Street, report.Patient.Address.Number, report.Patient.Address.City, report.Patient.Address.PostalCode)
		if err != nil {
			ErrorJSON(c, err.Error())
			return
		}

		report.Patient.Address.Address_id = int(*createdAddressID)

		createdPatientID, err = models.SavePatient(report.Patient)
		if err != nil {
			models.DeleteAdress(int(patient.Address.Address_id))
			ErrorJSON(c, err.Error())
			return
		}
	}

	fmt.Println("patient id =>", report.Patient.Patient_id)
	fmt.Println("created patient id =>", createdPatientID)

	// // create Report from models.SaveReport
	// newReportID, err := models.SaveReport(report.User_id, report.Patient_id, report.ReportContent, report.ArrivalTime, report.DepartureTime, report.AbscenceStatus)

	// if err != nil {
	// 	ErrorJSON(c, err.Error())
	// 	return
	// }

	c.JSON(200, gin.H{
		"ok":           true,
		"newPatientID": createdPatientID,
		"newAddressID": createdAddressID,
		"message":      "Report Created",
	})
}
