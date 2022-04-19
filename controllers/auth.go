package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	models.User
	jwt.StandardClaims
}

func NewToken(user models.User) (string, error) {

	expirationTime := time.Now().Add(120 * time.Hour) // 5 days

	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}

	return tokenString, nil
}

func ParseToken(token_string string, c *gin.Context) (*Claims, error) {

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Could not validate token",
		})

		c.Abort()
		return nil, nil
	}

	return claims, nil
}

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
	} else {
		var dbUser models.User
		stmt, err := utils.DB.Prepare("SELECT Email FROM Users WHERE Users.User_id = ? ;")
		if err != nil {
			fmt.Println("err while preparing stmt", err.Error())
		}

		stmt.QueryRow(user.User_id).Scan(&dbUser.Email)

		if dbUser.Email != "" {
			ErrorJSON(c, "User has already been assigned a mail")
			return
		}

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

	// TODO: kwdikos na apothikeyetai hashed
	// TODO: na epistrefei jwt token

	c.JSON(http.StatusOK, gin.H{
		"message":       "User updated successfully",
		"rows affected": rowsAffected,
	})

}

func Login(c *gin.Context) {
	fmt.Println("user logs in")
}
