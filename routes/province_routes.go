package routes

import (
	"net/http"

	database "wilayah_indonesia/config"
	"wilayah_indonesia/models"

	"github.com/gin-gonic/gin"
)

func GetProvinces(c *gin.Context) {
	var province []models.Province
	database.DB.Find(&province)

	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, province)

	//return json
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Lists Data Province",
		"data":    province,
	})
}
