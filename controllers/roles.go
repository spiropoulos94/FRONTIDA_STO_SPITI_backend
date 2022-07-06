package controllers

import (
	"net/http"
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/models"

	"github.com/gin-gonic/gin"
)

func ListRoles(c *gin.Context) {
	var roles []models.Profession
	roles, err := models.GetAllRoles()
	if err != nil {
		ErrorJSON(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"roles": roles,
		"ok":    true,
	})
}
