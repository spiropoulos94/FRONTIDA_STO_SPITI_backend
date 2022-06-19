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

func RetrieveDataFromContext(c *gin.Context) map[string]interface{} {
	m := make(map[string]interface{})
	userId, idExist := c.Get("User_id")
	if idExist {
		m["user_id"] = userId
	}
	userRoleId, roleIdExist := c.Get("User_Role_id")
	if roleIdExist {
		m["user_role_id"] = userRoleId
	}

	return m
}

func UserServices(c *gin.Context) {
	userData := RetrieveDataFromContext(c)

	userID := userData["user_id"].(int)
	userRoleID := userData["user_role_id"].(int)

	fmt.Println(userID, userRoleID)

	var services []models.Service
	var err error

	if userRoleID == 1 {
		fmt.Println("admin")
		services, err = models.GetAllServices()
		if err != nil {
			ErrorJSON(c, err.Error())
			return
		}

	} else {
		services, err = models.GetServicesByUserId(userID)
		if err != nil {
			ErrorJSON(c, err.Error())
			return
		}
		fmt.Println("user")
	}

	c.JSON(http.StatusOK, gin.H{"services": services})
}

func ListUsers(c *gin.Context) {
	fmt.Println("getting users!")

	contextData := RetrieveDataFromContext(c)

	fmt.Println(contextData)

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

func FindUser(c *gin.Context) {
	id := c.Param("id")

	var user *models.User

	user, err := models.GetUserByID(id)

	if err != nil {
		fmt.Println("err", err)
		ErrorJSON(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"users": user,
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

	newUserID, err := models.CreateUser(newUser.Name, newUser.Surname, newUser.AFM, newUser.AMKA, newUser.Profession.Role_id)

	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	newUserProfession, err := models.GetRole(newUser.Profession.Role_id)

	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	encodedFields := make(map[string]interface{})

	encodedFields["ID"] = newUserID
	encodedFields["Name"] = newUser.Name
	encodedFields["Surname"] = newUser.Surname
	encodedFields["AFM"] = newUser.AFM
	encodedFields["AMKA"] = newUser.AMKA
	encodedFields["Profession"] = newUserProfession

	fmt.Println(encodedFields)
	fmt.Println("--------")
	jsonByteSlice, err := json.Marshal(encodedFields)
	if err != nil {
		fmt.Println("err while marshal: ", err)
	}

	stringifiedJSON := string(jsonByteSlice)

	encodedStr := b64.StdEncoding.EncodeToString([]byte(stringifiedJSON))

	c.JSON(200, gin.H{
		"ok":                    true,
		"rows affected":         newUserID,
		"message":               "User added",
		"encodedFields(base64)": encodedStr,
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
