package routes

import (
	"net/http"

	database "wilayah_indonesia/config"
	"wilayah_indonesia/models"

	"github.com/gin-gonic/gin"
)

func GetRegencies(c *gin.Context) {
	var regencie []models.Regencie
	if err := database.DB.Where("province_id = ?", c.Param("province_id")).Find(&regencie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Regencies from province_id : " + c.Param("province_id"),
		"data":    regencie,
	})
}
