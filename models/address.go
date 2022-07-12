package models

import (
	"fmt"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
)

type Address struct {
	Address_id int64  `json:"Address_id"`
	Street     string `json:"Street"`
	Number     int    `json:"Number"`
	City       string `json:"City"`
	PostalCode int    `json:"Postal_code"`
}

func SaveAddress(street string, number int, city string, postalCode int) (*int64, error) {
	stmt, err := utils.DB.Prepare("INSERT INTO Addresses ( Street, Number, City, Postal_code) VALUES( ?, ?, ?, ? )")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(street, number, city, postalCode)
	if err != nil {
		return nil, err
	}

	newAdressID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	if numberOfRowsAffected, err := res.RowsAffected(); err != nil {
		fmt.Println("Rows affected,", numberOfRowsAffected)
		return nil, err
	} else {

		return &newAdressID, nil

	}
}

func DeleteAdress(addressID int) (bool, error) {
	stmt, err := utils.DB.Prepare("Delete from Addresses where Address_id = ?")
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(addressID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
