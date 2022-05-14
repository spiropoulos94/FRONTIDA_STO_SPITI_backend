package controllers

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"
	"strings"

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

	fmt.Println("admin create user runs")

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	newUser := models.User{}
	json.Unmarshal(jsonData, &newUser)

	if strings.TrimSpace(newUser.Name) == "" || strings.TrimSpace(newUser.Surname) == "" || newUser.AFM == 0 || newUser.AMKA == 0 || newUser.Profession.Role_id == 0 {
		ErrorJSON(c, "Name, Surname ,AFM and AMKA  and Role_id are needed ")
		return
	}

	rowsAffected, err := models.CreateUser(newUser.Name, newUser.Surname, newUser.AFM, newUser.AMKA, newUser.Profession.Role_id)

	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	encodedFields := make(map[string]interface{})

	encodedFields["Name"] = newUser.Name
	encodedFields["Surname"] = newUser.Surname
	encodedFields["AFM"] = newUser.AFM
	encodedFields["AMKA"] = newUser.AMKA
	encodedFields["Profession"] = newUser.Profession.Role_id

	fmt.Println(encodedFields)
	fmt.Println("--------")
	jsonByteSlice, err := json.Marshal(encodedFields)
	if err != nil {
		fmt.Println("err while marshal: ", err)
	}

	stringifiedJSON := string(jsonByteSlice)

	encodedStr := b64.StdEncoding.EncodeToString([]byte(stringifiedJSON))

	c.JSON(http.StatusOK, gin.H{
		"rows affected": rowsAffected,
		"message":       "User added",
		"encodedFields": encodedStr,
	})

	// stmt, err := utils.DB.Prepare("INSERT INTO Users( Name, Surname, AFM, AMKA, Role_id) VALUES( ?, ?, ?, ?, ? )")
	// if err != nil {
	// 	ErrorJSON(c, err.Error())
	// 	return
	// }

	// defer stmt.Close()

	// res, err := stmt.Exec(newUser.Name, newUser.Surname, newUser.AFM, newUser.AMKA, newUser.Profession.Role_id)

	// if err != nil {
	// 	ErrorJSON(c, err.Error())
	// 	return
	// }

	// if number, err := res.RowsAffected(); err != nil {
	// 	ErrorJSON(c, err.Error())
	// } else {

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"rows affected": number,
	// 		"message":       "User added",
	// 	})
	// }

}
