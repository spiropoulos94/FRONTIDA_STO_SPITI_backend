package models

import (
	"fmt"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
)

type Role struct {
	Role_id int    `json:"Role_id"`
	Title   string `json:"Title"`
}

func GetAllRoles() ([]Role, error) {
	var roles []Role

	rows, err := utils.DB.Query("SELECT Role_id, Title  FROM `Roles`;")
	if err != nil {
		fmt.Println("error => ", err)
		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var role Role

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

func GetRole(roleID int) (*Role, error) {
	var role Role

	if err := utils.DB.QueryRow("SELECT * from Roles where Role_id = ?", roleID).Scan(&role.Role_id, &role.Title); err != nil {
		return nil, err
	}

	return &role, nil
}
