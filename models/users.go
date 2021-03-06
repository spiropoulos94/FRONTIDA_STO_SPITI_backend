package models

import (
	"database/sql"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
)

type User struct {
	User_id    int        `json:"User_id"`
	Name       string     `json:"Name"`
	Surname    string     `json:"Surname"`
	AFM        int        `json:"AFM"`
	AMKA       int        `json:"AMKA"`
	Profession Profession `json:"Profession"`
	Email      string     `json:"Email"`
	Password   string     `json:"Password,omitempty"`

	Active       *bool `json:"Active,omitempty"`
	ReportsCount int   `json:"ReportsCount"`
}

type UserResponse struct {
	User_id    int        `json:"User_id"`
	Name       string     `json:"Name"`
	Surname    string     `json:"Surname"`
	AFM        int        `json:"AFM"`
	AMKA       int        `json:"AMKA"`
	Email      string     `json:"Email"`
	Profession Profession `json:"Profession"`
	Services   []Service  `json:"Services"`
	Reports    []Report   `json:"Reports,omitempty"`
}

func GenerateUserHash(id int) (*string, error) {

	user, err := GetUserByID(id)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	if len(user.Email) > 0 {
		return nil, errors.New("this user has already been registered")
	}

	encodedStr, err := UserToHashString(user)
	if err != nil {
		return nil, err
	}

	return encodedStr, nil
}

func GetUserByID(id int) (*User, error) {
	var user User

	stmt, err := utils.DB.Prepare("SELECT User_id, Name, Surname, AFM, AMKA, Email, Roles.Role_id, Roles.Title FROM Users LEFT JOIN Roles ON Users.Role_id = Roles.Role_id WHERE User_id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var userEmail sql.NullString

	err = row.Scan(&user.User_id, &user.Name, &user.Surname, &user.AFM, &user.AMKA, &userEmail, &user.Profession.Role_id, &user.Profession.Title)

	if err != nil {
		return nil, err
	}

	user.Email = userEmail.String

	return &user, nil

}

func GetAllUsers() ([]User, error) {
	var users []User

	rows, err := utils.DB.Query("SELECT Users.User_id, Users.Name, Users.Surname, Users.Email, Users.AFM, Users.AMKA,  Roles.Title , Roles.Role_id  FROM `Users` left join Roles on users.Role_id = Roles.Role_id")

	if err != nil {
		fmt.Println("error => ", err)
		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user User
		var profession Profession
		// var report Report

		var userEmail sql.NullString

		if err := rows.Scan(&user.User_id, &user.Name, &user.Surname, &userEmail, &user.AFM, &user.AMKA, &profession.Title, &profession.Role_id); err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		user.Profession = profession
		user.Email = userEmail.String
		userActiveStatus := len(userEmail.String) > 0
		user.Active = &userActiveStatus

		userReportsCount, err := GetReportsCount(user.User_id)
		if err != nil {
			return nil, err
		}

		user.ReportsCount = userReportsCount

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return users, nil
}

func CreateUser(name, surname string, AFM, AMKA, roleID int) (int64, error) {
	stmt, err := utils.DB.Prepare("INSERT INTO Users( Name, Surname, AFM, AMKA, Role_id) VALUES( ?, ?, ?, ?, ? )")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(name, surname, AFM, AMKA, roleID)
	if err != nil {
		return 0, err
	}

	newUserID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	if err != nil {
		return 0, err
	}

	if numberOfRowsAffected, err := res.RowsAffected(); err != nil {
		fmt.Println("Rows affected,", numberOfRowsAffected)
		return 0, err
	} else {

		return newUserID, nil

	}

}

func UserToHashString(user *User) (*string, error) {
	encodedFields := make(map[string]interface{})

	encodedFields["ID"] = user.User_id
	encodedFields["Name"] = user.Name
	encodedFields["Surname"] = user.Surname
	encodedFields["AFM"] = user.AFM
	encodedFields["AMKA"] = user.AMKA
	encodedFields["Profession"] = user.Profession

	jsonByteSlice, err := json.Marshal(encodedFields)
	if err != nil {
		return nil, err
	}

	stringifiedJSON := string(jsonByteSlice)

	encodedStr := b64.StdEncoding.EncodeToString([]byte(stringifiedJSON))

	return &encodedStr, nil
}
