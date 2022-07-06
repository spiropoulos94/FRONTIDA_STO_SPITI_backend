package models

import (
	"fmt"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
)

type Profession struct {
	Role_id int    `json:"Role_id"`
	Title   string `json:"Title"`
}

func GetAllRoles() ([]Profession, error) {
	var roles []Profession

	rows, err := utils.DB.Query("SELECT Role_id, Title  FROM `Roles`;")
	if err != nil {
		fmt.Println("error => ", err)
		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var role Profession

		if err := rows.Scan(&role.Role_id, &role.Title); err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return roles, nil
}

func GetRole(professionID int) (*Profession, error) {
	var role Profession

	if err := utils.DB.QueryRow("SELECT * from Roles where Role_id = ?", professionID).Scan(&role.Role_id, &role.Title); err != nil {
		return nil, err
	}

	return &role, nil
}
