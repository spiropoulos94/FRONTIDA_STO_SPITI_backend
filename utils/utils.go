package utils

import (
	"fmt"
)

//d arithmos, s string

func ID_exists(idColumn string, idVal int, table string) (bool, error) {

	query := fmt.Sprintf("SELECT %s.%s FROM %s WHERE %s.%s = %d ;", table, idVal, table, table, idColumn, idVal)

	fmt.Println("query =>", query)

	stmt, err := DB.Prepare(query)

	if err != nil {
		fmt.Println("error in preparing id_exists statement")
		return false, err
	}

	defer stmt.Close()

}
