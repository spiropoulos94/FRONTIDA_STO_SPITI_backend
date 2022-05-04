package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/controllers"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckHeaderForJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MIDDLEWARE checking for authorization headers")

		authoriazationHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authoriazationHeader, "Bearer ") {
			controllers.ErrorJSON(c, "You are not authorized")
			c.Abort()
			return
		}

		token := strings.Split(authoriazationHeader, " ")[1]

		user, err := controllers.ParseToken(token, c)

		fmt.Println("parse token and get this => ", user)

		if err != nil {
			controllers.ErrorJSON(c, err.Error())
			c.Abort()
			return
		}

		dbStoredUser := models.User{}

		stmt, err := utils.DB.Prepare("SELECT Users.User_id, Users.Role_id FROM Users LEFT JOIN Roles ON Users.Role_id = Roles.Role_id WHERE Users.User_id = ? ;")

		if err != nil {
			fmt.Println(2)

			controllers.ErrorJSON(c, err.Error())
			c.Abort()
			return
		}

		defer stmt.Close()

		err = stmt.QueryRow(user.User_id).Scan(&dbStoredUser.User_id, &dbStoredUser.Profession.Role_id)

		if err == sql.ErrNoRows || err != nil {
			fmt.Println("Error", err)
			controllers.ErrorJSON(c, err.Error())
			c.Abort()
			return
		}

		c.Set("User_id", dbStoredUser.User_id)
		c.Set("User_Role_id", dbStoredUser.Profession.Role_id)

		c.Next()
	}
}

func CheckIfUserIsAdmin() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("MIDDLEWARE checking for if user is admin in order to proceed")

		roleId, exists := c.Get("User_Role_id")

		if exists && roleId == 1 {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "No access",
			})
			c.Abort()
		}
	}

}
