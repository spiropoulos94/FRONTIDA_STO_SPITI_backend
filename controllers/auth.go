package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	fmt.Println("sign up!!")
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	newUser := models.User{}
	json.Unmarshal(jsonData, &newUser)

	fmt.Println(newUser.User_id)

	stmt, err := utils.DB.Prepare("SELECT * FROM Users WHERE Users.User_id = ?")

	if err != nil {
		ErrorJSON(c, err)
		return
	}

	defer stmt.Close()

}
