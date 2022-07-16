package controllers

import (
	"net/http"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"

	"github.com/gin-gonic/gin"
)

func ListPatients(c *gin.Context) {

	var patients []models.Patient
	patients, err := models.GetAllPatients()
	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"patients": patients,
		"ok":       true,
	})
}
