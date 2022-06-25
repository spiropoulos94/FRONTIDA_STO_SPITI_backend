package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
	"strconv"
	"strings"
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
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}

	return tokenString, nil
}

func ParseToken(token_string string, c *gin.Context) (*Claims, error) {

	claims := &Claims{}
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	_, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {

		return nil, err
	}

	return claims, nil
}

// complete sign up
func CompleteSignUp(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	reqBodyUser := models.User{}
	json.Unmarshal(jsonData, &reqBodyUser)

	if reqBodyUser.Email == "" || reqBodyUser.Password == "" || reqBodyUser.User_id == 0 {
		ErrorJSON(c, "Email and Password & User_id are needed")
		return
	}

	userExists, err := utils.IDexistsInTable("Users", "User_id", reqBodyUser.User_id)

	if err != nil {
		fmt.Println("Err while checking if id exists")
		ErrorJSON(c, err.Error())
		return
	}

	var dbUser models.User

	if !userExists {
		ErrorJSON(c, "User does not exist in db")
		return
	} else {
		stmt, err := utils.DB.Prepare("SELECT Email FROM Users WHERE Users.User_id = ? ;")
		if err != nil {
			fmt.Println("err while preparing stmt", err.Error())
		}

		defer stmt.Close()

		stmt.QueryRow(reqBodyUser.User_id).Scan(&dbUser.Email)

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

	hashedPassword, err := utils.HashPassword(reqBodyUser.Password)

	if err != nil {
		fmt.Println("Error while hashing password")
		ErrorJSON(c, err.Error())
		return
	}

	res, err := updateStmt.Exec(reqBodyUser.Email, hashedPassword, reqBodyUser.User_id)

	if err != nil {
		fmt.Println("error while executing statement")
		ErrorJSON(c, err.Error())
		return
	}

	rowsAffected, _ := res.RowsAffected()

	stmt, err := utils.DB.Prepare("SELECT User_id, Role_id FROM Users WHERE User_id = ? ;")

	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	err = stmt.QueryRow(reqBodyUser.User_id).Scan(&dbUser.User_id, &dbUser.Profession.Role_id)

	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	token, err := NewToken(dbUser)

	if err != nil {
		fmt.Println("error while making token")
		ErrorJSON(c, err.Error())
		return
	}

	createdUserFromDB, err := models.GetUserByID(strconv.Itoa(dbUser.User_id))
	if err != nil {
		fmt.Println("error while making token")
		ErrorJSON(c, err.Error())
		return
	}

	var availableServices []models.Service
	var availableRoles []models.Role

	if dbUser.Profession.Role_id == 1 {
		availableServices, err = models.GetAllServices()
		if err != nil {
			ErrorJSON(c, err.Error())
			return
		}
		availableRoles, err = models.GetAllRoles()
		if err != nil {
			ErrorJSON(c, err.Error())
			return
		}
	} else {
		availableServices, err = models.GetServicesByUserId(dbUser.User_id)
		if err != nil {
			ErrorJSON(c, err.Error())
			return
		}
	}

	userResponse := models.UserLoginReponse{
		User_id:    createdUserFromDB.User_id,
		Name:       createdUserFromDB.Name,
		Surname:    createdUserFromDB.Surname,
		AFM:        createdUserFromDB.AFM,
		AMKA:       createdUserFromDB.AMKA,
		Email:      createdUserFromDB.Email,
		Profession: createdUserFromDB.Profession,
		Services:   availableServices,
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":            true,
		"message":       "User signup complete!",
		"rows affected": rowsAffected,
		"token":         token,
		"user":          userResponse,
		"roles":         availableRoles,
	})

}

func Login(c *gin.Context) {

	fmt.Println("Login")

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	reqBodyUser := models.User{}
	json.Unmarshal(jsonData, &reqBodyUser)

	fmt.Println("1")

	// for development purposes, admin always passes and get jwt
	// environment := os.Getenv("ENVIRONMENT")
	// if reqBodyUser.Email == "dev@dev.gr" && environment == "dev" {
	// 	fmt.Println("dev logged in ")

	// 	token, err := NewToken(reqBodyUser)

	// 	if err != nil {
	// 		ErrorJSON(c, err.Error())
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "DEV logged in successfuly",
	// 		"token":   token,
	// 	})

	// 	return
	// }

	if strings.TrimSpace(reqBodyUser.Email) == "" || strings.TrimSpace(reqBodyUser.Password) == "" {
		ErrorJSON(c, "Mail and password are needed")
		return
	}

	stmt, err := utils.DB.Prepare("Select Name, Surname, Email, Password, User_id, Roles.Role_id, Roles.Title from Users LEFT JOIN Roles ON Roles.Role_id = Users.Role_id WHERE Email = ? ")
	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	defer stmt.Close()

	var dbUser models.User
	var dbUserProfession models.Profession

	err = stmt.QueryRow(reqBodyUser.Email).Scan(&dbUser.Name, &dbUser.Surname, &dbUser.Email, &dbUser.Password, &dbUser.User_id, &dbUserProfession.Role_id, &dbUserProfession.Title)

	dbUser.Profession = dbUserProfession

	if err == sql.ErrNoRows {
		fmt.Println("No Rows for id", err)
		ErrorJSON(c, "User does not exist")
		return

	} else if err != nil {
		// edw skaei
		fmt.Println("Error", err)
		ErrorJSON(c, err.Error())
		return
	}

	if !utils.CheckPasswordHash(reqBodyUser.Password, dbUser.Password) {
		fmt.Println("wrong password")
		ErrorJSON(c, "Wrong password")
		return
	} else {
		fmt.Println("password ok")

		dbUser.Password = ""

		token, err := NewToken(dbUser)

		if err != nil {
			ErrorJSON(c, err.Error())
			return
		}

		var availableServices []models.Service
		var availableRoles []models.Role

		if dbUser.Profession.Role_id == 1 {
			availableServices, err = models.GetAllServices()
			availableRoles, err = models.GetAllRoles()
		} else {
			availableServices, err = models.GetServicesByUserId(dbUser.User_id)
		}

		if err != nil {
			ErrorJSON(c, err.Error())
			return
		}

		userResponse := models.UserLoginReponse{
			User_id:    dbUser.User_id,
			Name:       dbUser.Name,
			Surname:    dbUser.Surname,
			AFM:        dbUser.AFM,
			AMKA:       dbUser.AMKA,
			Email:      dbUser.Email,
			Profession: dbUser.Profession,
			Services:   availableServices,
		}

		c.JSON(http.StatusOK, gin.H{
			"ok":       true,
			"message":  "user logged in successfuly",
			"token":    token,
			"services": availableServices,
			"roles":    availableRoles,
			"user":     userResponse, // TODO: ftiakse user response struct pou na exei mesa ID, Name, Surname, Mail, Profession, Services
		})
	}

}
