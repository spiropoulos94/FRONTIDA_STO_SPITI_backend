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

func InsertAdminAccount() error {
	fmt.Println("insert admin account runs")
	//If admin account does not exist, make one.
	query := "SELECT Users.User_id from Users WHERE Email = 'dev@dev.gr';"

	var id int

	err := DB.QueryRow(query).Scan(&id)

	if err == sql.ErrNoRows {
		fmt.Println("No admin account found, will create one.", err)

		hashedAdminPass, err := HashPassword(os.Getenv("ADMIN_PASS"))

		if err != nil {
			fmt.Println("Error while hashing password")
			return err
		}

		createAdminQuery := "INSERT INTO Users ( Name, Surname, AFM, AMKA, Role_id,  Email, Password) VALUES ('AdminDev', 'AdminDev', '0' , '0', '1', 'dev@dev.gr', ?);"

		res, err := DB.Exec(createAdminQuery, hashedAdminPass)

		if err != nil {
			return err
		}

		if numberOfRowsAffected, err := res.RowsAffected(); err != nil {
			return err
		} else {

			fmt.Println("Number of rows affected : ", numberOfRowsAffected)
			return nil

		}

	} else if err != nil {
		fmt.Println("Error : ", err)
		return err
	} else {
		fmt.Println("admin account found in database")
		return nil
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

	createAdminErr := InsertAdminAccount()

	if createAdminErr != nil {
		fmt.Println("!!! Error while creating admin account : ", err)
		log.Fatal(createAdminErr)
	}

}
