package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	fmt.Println("sign up!!")
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	user := models.User{}
	json.Unmarshal(jsonData, &user)

	fmt.Println(user.User_id)

	stmt, err := utils.DB.Prepare("SELECT Users.User_id FROM Users WHERE Users.User_id = ?")

	if err != nil {
		fmt.Println("Error in preparing statement")
		ErrorJSON(c, err.Error())
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(user.User_id).Scan(&user.User_id)

	if err == sql.ErrNoRows {
		ErrorJSON(c, "User doesn't exist")
		return
	} else if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	fmt.Println("User exists!!")

	updateStmt, err := utils.DB.Prepare("UPDATE Users SET Email = ?, Password = ?,	WHERE User_id = ?")

	if err != nil {
		fmt.Println("error preparing the update statement")
		ErrorJSON(c, err.Error())
		return
	}

	defer updateStmt.Close()

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
