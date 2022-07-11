package models

import (
	"database/sql"
	"errors"
	"fmt"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
)

type Patient struct {
	Patient_id     int     `json:"Patient_id"`
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
			err = errors.New("404")
		}
		return nil, err
	}

	return &patient, nil

}

func SavePatient(patient Patient) (int64, error) {
	// create patient address before creating Patient

	fmt.Println("* Begining to save patient")

	addressID, err := SaveAddress(patient.Address.Street, patient.Address.Number, patient.Address.City, patient.Address.PostalCode)
	if err != nil {
		return 0, err
	}
	fmt.Println("* Patient address saved succesfully! ")

	stmt, err := utils.DB.Prepare("INSERT INTO Patients ( Fullname, Patient_AMKA, Health_security, Address_id) VALUES( ?, ?, ?, ? )")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(patient.Fullname, patient.Patient_AMKA, patient.HealthSecurity, addressID)
	if err != nil {
		// if patient fails to be saved, erase his address to prevent bloated database
		fmt.Println("failed to create patient")
		_, err := DeleteAdress(int(addressID))
		if err != nil {
			fmt.Println("Failed to delte address")
			return -1, err
		}

	}

	newPatientID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	if numberOfRowsAffected, err := res.RowsAffected(); err != nil {
		fmt.Println("Rows affected,", numberOfRowsAffected)
		return 0, err
	} else {
		return newPatientID, nil
	}
}
