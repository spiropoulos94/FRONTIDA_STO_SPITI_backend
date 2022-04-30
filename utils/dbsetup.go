package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InsertAdminAccount() {
	fmt.Println("insert admin account runs")
	//delete admin account if exists and make a new one
	query := "SELECT Users.User_id from Users WHERE Email = 'dev@dev.gr';"

	var id int

	err := DB.QueryRow(query).Scan(&id)

	if err == sql.ErrNoRows {
		fmt.Println("No admin account found", err)
		return
	} else if err != nil {
		fmt.Println("Error : ", err)
		return
	} else {
		fmt.Println("admin account found in database")
		return
	}

}

func SetupDatabase() {

	godotenv.Load(".env")

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	InsertAdminAccount()

}
