package routes

import (
	"net/http"

	database "wilayah_indonesia/config"
	"wilayah_indonesia/models"

	"github.com/gin-gonic/gin"
)

func GetVilanges(c *gin.Context) {
	var villages []models.Villages
	if err := database.DB.Where("district_id = ?", c.Param("district_id")).Find(&villages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "villagess from district_id : " + c.Param("district_id"),
		"data":    villages,
	})
}
