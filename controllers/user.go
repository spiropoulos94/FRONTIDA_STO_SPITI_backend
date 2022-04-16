package controllers

import (
	"fmt"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"

	"github.com/gin-gonic/gin"
)

func ErrorJSON(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"users": err,
	})
}

func ListUsers(c *gin.Context) {
	fmt.Println("getting users!")

	var users []models.User

	rows, err := utils.DB.Query("SELECT Users.User_id, Users.Name, Users.Surname, Users.AFM, Users.AMKA, Users.Email, Users.Password,  Roles.Title as Profession  FROM `Users` left join Roles on users.Role_id = Roles.Role_id")
	if err != nil {
		fmt.Println("error => ", err)
		ErrorJSON(c, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.User_id, &user.Name, &user.Surname, &user.AFM, &user.AMKA, &user.Email, &user.Password, &user.Profession); err != nil {
			fmt.Println("err", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("err", err)
		ErrorJSON(c, err)
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}
