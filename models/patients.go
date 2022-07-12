package models

import (
	"database/sql"
	"errors"
	"fmt"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
)

type Patient struct {
	Patient_id     int64   `json:"Patient_id"`
	Fullname       string  `json:"Fullname"`
	Patient_AMKA   int     `json:"Patient_AMKA"`
	HealthSecurity bool    `json:"Health_security"`
	Address        Address `json:"Address"`
}

func GetPatientByAMKA(amka int) (*Patient, error) {
	var patient Patient
	stmt, err := utils.DB.Prepare("SELECT Patient_id, Fullname, Patient_AMKA, Health_security, Street, Number, City, Postal_code FROM `Patients` left join Addresses on Patients.Address_id = Addresses.Address_id WHERE Patients.Patient_AMKA = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(amka)

	err = row.Scan(&patient.Patient_id, &patient.Fullname, &patient.Patient_AMKA, &patient.HealthSecurity, &patient.Address.Street, &patient.Address.Number, &patient.Address.City, &patient.Address.PostalCode)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("404 patient does not exist")
		}
		return nil, err
	}

	return &patient, nil

}

func SavePatient(patient Patient) (*int64, error) {
	// create patient address before creating Patient

	stmt, err := utils.DB.Prepare("INSERT INTO Patients ( Fullname, Patient_AMKA, Health_security, Address_id) VALUES( ?, ?, ?, ? )")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(patient.Fullname, patient.Patient_AMKA, patient.HealthSecurity, patient.Address.Address_id)
	if err != nil {
		if err != nil {
			return nil, err
		}

	}

	newPatientID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	if numberOfRowsAffected, err := res.RowsAffected(); err != nil {
		fmt.Println("Rows affected,", numberOfRowsAffected)
		return nil, err
	} else {
		return &newPatientID, nil
	}
}
