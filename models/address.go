package models

import (
	"fmt"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
)

type Address struct {
	Address_id int    `json:"Address_id"`
	Street     string `json:"Street"`
	Number     int    `json:"Number"`
	City       string `json:"City"`
	PostalCode int    `json:"Postal_code"`
}

func SaveAddress(street string, number int, city string, postalCode int) (int64, error) {
	stmt, err := utils.DB.Prepare("INSERT INTO Addresses ( Street, Number, City, Postal_code) VALUES( ?, ?, ?, ? )")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(street, number, city, postalCode)
	if err != nil {
		return -1, err
	}

	newAdressID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	if numberOfRowsAffected, err := res.RowsAffected(); err != nil {
		fmt.Println("Rows affected,", numberOfRowsAffected)
		return 0, err
	} else {

		return newAdressID, nil

	}
}
