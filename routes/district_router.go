package routes

import (
	"net/http"

	database "wilayah_indonesia/config"
	"wilayah_indonesia/models"

	"github.com/gin-gonic/gin"
)

func GetDistrict(c *gin.Context) {
	var district []models.District
	if err := database.DB.Where("regencie_id = ?", c.Param("regencie_id")).Find(&district).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "districts from regencie_id : " + c.Param("regencie_id"),
		"data":    district,
	})
}
