package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"

	"github.com/gin-gonic/gin"
)

func ErrorJSON(c *gin.Context, err interface{}) {
	c.JSON(http.StatusForbidden, gin.H{
		"error": err,
	})
}

func ListUsers(c *gin.Context) {
	fmt.Println("getting users!")

	userId, exist := c.Get("User_id")

	if exist {
		fmt.Println("User id", userId)
	}
	userRoleId, exist := c.Get("User_Role_id")

	if exist {
		fmt.Println("Role id", userRoleId)
	}

	var users []models.User

	users, err := models.GetAllUsers()

	if err != nil {
		fmt.Println("err", err)
		ErrorJSON(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func AdminCreateUser(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	newUser := models.User{}
	json.Unmarshal(jsonData, &newUser)

	if newUser.AFM == 0 || newUser.AMKA == 0 || newUser.Profession.Role_id == 0 {
		ErrorJSON(c, "AFM and AMKA  and Role_id is needed ")
		return
	}

	stmt, err := utils.DB.Prepare("INSERT INTO Users( Name, Surname, AFM, AMKA, Role_id) VALUES( ?, ?, ?, ?, ? )")
	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	defer stmt.Close()

	res, err := stmt.Exec(newUser.Name, newUser.Surname, newUser.AFM, newUser.AMKA, newUser.Profession.Role_id)

	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	if number, err := res.RowsAffected(); err != nil {
		ErrorJSON(c, err.Error())
	} else {

		c.JSON(http.StatusOK, gin.H{
			"rows affected": number,
			"message":       "User added",
		})
	}

}
