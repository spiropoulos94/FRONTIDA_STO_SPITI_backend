package models

import (
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
	Password   string     `json:"Password"`
}

type Profession struct {
	Role_id int    `json:"Role_id"`
	Title   string `json:"Title"`
}

func GetAllUsers() ([]User, error) {
	var users []User

	rows, err := utils.DB.Query("SELECT Users.User_id, Users.Name, Users.Surname, Users.AFM, Users.AMKA,  Roles.Title , Roles.Role_id  FROM `Users` left join Roles on users.Role_id = Roles.Role_id")
	if err != nil {
		fmt.Println("error => ", err)
		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user User
		var profession Profession

		if err := rows.Scan(&user.User_id, &user.Name, &user.Surname, &user.AFM, &user.AMKA, &profession.Title, &profession.Role_id); err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		user.Profession = profession
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

	if numberOfRowsAffected, err := res.RowsAffected(); err != nil {
		return 0, err
	} else {

		return numberOfRowsAffected, nil

	}

}
