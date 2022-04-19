package utils

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
