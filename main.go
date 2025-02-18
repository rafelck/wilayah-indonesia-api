package main

import (
	database "wilayah_indonesia/config"
	"wilayah_indonesia/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	database.ConnectDatabase()

	// Setup router
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "It's free",
		})
	})
	r.GET("/provinces", routes.GetProvinces)
	r.GET("/regencies/:province_id", routes.GetRegencies)
	r.GET("/district/:regencie_id", routes.GetDistrict)
	r.GET("/villages/:district_id", routes.GetVilanges)

	// Jalankan server
	r.Run(":8080")
}
