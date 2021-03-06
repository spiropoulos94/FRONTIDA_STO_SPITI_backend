package models

import (
	"fmt"
	"log"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
	"strconv"
)

type Service struct {
	ID      int    `json:"Service_id"`
	Title   string `json:"Title"`
	Role_id int    `json:"Role_id"`
}

func GetAllServices() ([]Service, error) {
	var services []Service

	rows, err := utils.DB.Query("SELECT Service_id, Title, Role_id FROM `Services` ")
	if err != nil {
		fmt.Println("error => ", err)
		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var service Service

		if err := rows.Scan(&service.ID, &service.Title, &service.Role_id); err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return services, nil
}

func GetServicesByRoleId(id int) ([]Service, error) {
	var services []Service

	fmt.Println("getting services for user with id: ", id)

	stmt, err := utils.DB.Prepare("SELECT Service_id, Title, Role_id FROM `Services`  WHERE Role_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Println("error => ", err)
		return nil, err
	}

	for rows.Next() {
		var service Service

		if err := rows.Scan(&service.ID, &service.Title, &service.Role_id); err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		services = append(services, service)

	}
	if err := rows.Err(); err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return services, nil

}

func SaveReportServices(reportID int, servicesIDs []int) (*int64, error) {

	queryString := "INSERT INTO Reports_services ( Report_id, Service_id ) VALUES"

	values := make([]interface{}, 0, len(servicesIDs)*2)

	reportIDStr := strconv.Itoa(reportID)

	for i, serviceID := range servicesIDs {

		connectString := ","
		if i == (len(servicesIDs) - 1) {
			connectString = ";"
		}

		values = append(values, reportIDStr, serviceID)

		valuesString := fmt.Sprintf("(?, ?)%s", connectString)

		queryString = queryString + valuesString
	}

	stmt, err := utils.DB.Prepare(queryString)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(values...)
	if err != nil {
		return nil, err
	}

	numberOfRowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &numberOfRowsAffected, nil
}
