package models

import (
	"fmt"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
)

type Report struct {
	Report_id      int     `json:"Report_id"`         // ok
	Author         Author  `json:"Author"`            // ase ayta gia to telos
	Patient        Patient `json:"Patient"`           // ase ayta gia to telos
	ReportContent  string  `json:"Report_content"`    // ok
	ReportDate     int     `json:"Report_Date_ts"`    // ok
	ArrivalTime    int     `json:"Arrival_Time_ts"`   // ok
	DepartureTime  int     `json:"Departure_Time_ts"` // ok
	AbscenceStatus bool    `json:"Absence_Status"`    // ok
}

type UserReportResponse struct {
	Report_id      int     `json:"Report_id"`
	Patient        Patient `json:"Patient"`
	ReportContent  string  `json:"Report_content"`    // ok
	ReportDate     int     `json:"Report_Date_ts"`    // ok
	ArrivalTime    int     `json:"Arrival_Time_ts"`   // ok
	DepartureTime  int     `json:"Departure_Time_ts"` // ok
	AbscenceStatus bool    `json:"Absence_Status"`
}

type Author struct {
	User_id    int        `json:"User_id"`    // ok
	Name       string     `json:"Name"`       // ok
	Surname    string     `json:"Surname"`    // ok
	Profession Profession `json:"Profession"` // ok
}

type Patient struct {
	Patient_id     int     `json:"Patient_id"`
	Fullname       string  `json:"Fullname"`
	Patient_AMKA   int     `json:"Patient_AMKA"`
	HealthSecurity bool    `json:"Health_security"`
	Address        Address `json:"Address"`
}

type Address struct {
	Address_id int    `json:"Address_id"`
	Street     string `json:"Street"`
	Number     int    `json:"Number"`
	City       string `json:"City"`
	PostalCode int    `json:"Postal_code"`
}

func GetAllReports() ([]Report, error) {

	var reports []Report

	// se ayto to query enwse reports, users kai patients gia na pareis ta data
	stmt, err := utils.DB.Prepare("SELECT Daily_Reports.Report_id, Daily_Reports.Report_content, Daily_Reports.Report_Date_ts, Daily_Reports.Arrival_Time_ts, Daily_Reports.Departure_Time_ts, Daily_Reports.Absence_Status, Users.User_id, Users.Name, Users.Surname, Roles.Role_id, Roles.Title, Patients.Patient_id, Patients.Fullname, Patients.Patient_AMKA, Patients.Health_security, Addresses.Address_id, Addresses.Street, Addresses.Number, Addresses.City, Addresses.Postal_code   FROM Daily_Reports LEFT JOIN Roles ON Daily_Reports.User_id = Roles.Role_id	LEFT JOIN Users ON Daily_Reports.User_id = Users.User_id LEFT JOIN Patients ON Daily_Reports.Patient_id = Patients.Patient_id LEFT JOIN Addresses ON Patients.Address_id = Addresses.Address_id ;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var report Report
		var author Author
		var patient Patient
		var profession Profession
		var address Address

		if err := rows.Scan(&report.Report_id, &report.ReportContent, &report.ReportDate, &report.ArrivalTime, &report.DepartureTime, &report.AbscenceStatus, &author.User_id, &author.Name, &author.Surname, &profession.Role_id, &profession.Title, &patient.Patient_id, &patient.Fullname, &patient.Patient_AMKA, &patient.HealthSecurity, &address.Address_id, &address.Street, &address.Number, &address.City, &address.PostalCode); err != nil {
			fmt.Println("err", err)
			return nil, err
		}

		author.Profession = profession
		report.Author = author
		patient.Address = address
		report.Patient = patient

		reports = append(reports, report)
	}

	return reports, nil
}

func GetReportsCount(userID int) (int, error) {
	stmt, err := utils.DB.Prepare("SELECT  count(*) as ReportsCount FROM `Users` left join Daily_Reports on Daily_Reports.User_id = users.User_id 	WHERE Daily_Reports.User_id IS NOT NULL AND Users.User_id = ?")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	var rowsNumber int

	err = stmt.QueryRow(userID).Scan(&rowsNumber)
	if err != nil {
		return -1, err
	}

	return rowsNumber, nil
}

func GetUserReports(userID int) ([]UserReportResponse, error) {
	// var userReports []Report
	userReports := make([]UserReportResponse, 0)
	stmt, err := utils.DB.Prepare("SELECT Daily_Reports.Report_id, Daily_Reports.Report_content, Report_Date_ts, Arrival_Time_ts, Departure_Time_ts, Daily_Reports.Absence_Status, Patients.Patient_id, Patients.Fullname as Patient_Fullname, Patients.Health_security, Addresses.Street, Addresses.Number, Addresses.City, Addresses.Postal_code FROM `Daily_Reports` LEFT JOIN Patients ON Daily_Reports.Patient_id = Patients.Patient_id LEFT JOIN Addresses on Patients.Address_id = Addresses.Address_id WHERE Daily_Reports.User_id =  ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var report UserReportResponse
		var patient Patient
		var address Address

		if err := rows.Scan(&report.Report_id, &report.ReportContent, &report.ReportDate, &report.ArrivalTime, &report.DepartureTime, &report.AbscenceStatus, &patient.Patient_id, &patient.Fullname, &patient.HealthSecurity, &address.Street, &address.Number, &address.City, &address.PostalCode); err != nil {
			fmt.Println("err", err)
			return nil, err
		}

		patient.Address = address
		report.Patient = patient

		userReports = append(userReports, report)

	}

	return userReports, nil
}
