package controllers

import (
	"database/sql"
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

	var patientID *int64
	var addressID *int64

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	report := models.Report{}
	json.Unmarshal(jsonData, &report)

	if report.Patient.Patient_AMKA == 0 || report.Patient.Patient_AMKA > 999999999 {
		ErrorJSON(c, "Use a valid AMKA")
		return
	}

	// check if dbPatient already exists
	dbPatient, err := models.GetPatientByAMKA(report.Patient.Patient_AMKA)

	if err != nil && err != sql.ErrNoRows {
		ErrorJSON(c, err.Error())
		return
	}

	if dbPatient == nil {
		addressID, err := models.SaveAddress(report.Patient.Address.Street, report.Patient.Address.Number, report.Patient.Address.City, report.Patient.Address.PostalCode)
		if err != nil {
			ErrorJSON(c, err.Error())
			return
		}

		report.Patient.Address.Address_id = *addressID

		patientID, err = models.SavePatient(report.Patient)
		if err != nil {
			models.DeleteAdress(int(dbPatient.Address.Address_id))
			ErrorJSON(c, err.Error())
			return
		}
	} else {
		patientID = &dbPatient.Patient_id
		addressID = &dbPatient.Address.Address_id
	}

	if patientID != nil {
		fmt.Println("dbPatient id =>", *patientID)
		fmt.Println("dbPatient address id =>", *addressID)
		fmt.Println("dbPatient id =>", *patientID)
		fmt.Println("dbPatient address id =>", *addressID)
		fmt.Println("dbPatient id =>", *patientID)
		fmt.Println("dbPatient address id =>", *addressID)
	}

	// // create Report from models.SaveReport
	// newReportID, err := models.SaveReport(report.User_id, report.Patient_id, report.ReportContent, report.ArrivalTime, report.DepartureTime, report.AbscenceStatus)

	// if err != nil {
	// 	ErrorJSON(c, err.Error())
	// 	return
	// }

	c.JSON(200, gin.H{
		"ok":        true,
		"PatientID": patientID,
		"AddressID": addressID,
		"message":   "Report Created",
	})
}
