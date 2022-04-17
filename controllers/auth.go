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

func SignUp(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	user := models.User{}
	json.Unmarshal(jsonData, &user)

	userExists, err := utils.IDexistsInTable("Users", "User_id", user.User_id)

	if err != nil {
		fmt.Println("Err while checking if id exists")
		ErrorJSON(c, err.Error())
	}

	if !userExists {
		ErrorJSON(c, "User does not exist in db")
		return
	}
	fmt.Println("mail", user.Email)
	if user.Email != "" {
		ErrorJSON(c, "Cannot complete signup, user already exists")
		return
	}

	updateStmt, err := utils.DB.Prepare("UPDATE Users SET Email = ?, Password = ? WHERE User_id = ?;")

	if err != nil {
		fmt.Println("error preparing the update statement")
		ErrorJSON(c, err.Error())
		return
	}

	defer updateStmt.Close()

	if user.Email == "" || user.Password == "" {
		ErrorJSON(c, "Mail and password are needed")
		return
	}
	res, err := updateStmt.Exec(user.Email, user.Password, user.User_id)

	if err != nil {
		fmt.Println("error while executing statement")
		ErrorJSON(c, err.Error())
		return
	}

	rowsAffected, _ := res.RowsAffected()

	c.JSON(http.StatusOK, gin.H{
		"message":       "User updated successfully",
		"rows affected": rowsAffected,
	})

}
