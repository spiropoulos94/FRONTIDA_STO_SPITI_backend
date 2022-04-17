package utils

import (
	"database/sql"
	"fmt"
)

//d arithmos, s string

func IDexistsInTable(table string, idColumn string, idVal int) (bool, error) {

	query := fmt.Sprintf("SELECT %s.%s FROM %s WHERE %s.%s = ? ;", table, idColumn, table, table, idColumn)

	fmt.Println("query =>", query)

	stmt, err := DB.Prepare(query)

	if err != nil {
		fmt.Println("error in preparing id_exists statement")
		return false, err
	}

	defer stmt.Close()

	var id int
	err = stmt.QueryRow(idVal).Scan(&id)

	if err == sql.ErrNoRows {
		fmt.Println("No Rows for id", err)
		return false, nil
	} else if err != nil {
		fmt.Println("Error", err)
		return false, err
	} else {
		return true, nil
	}

}
