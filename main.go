package main

import (
	"github.com/joho/godotenv"
	"os"
	database "wilayah_indonesia/config"
	"wilayah_indonesia/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	database.ConnectDatabase()
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	greeting := os.Getenv("GREETING")
	// Setup router
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": greeting,
		})
	})
	r.GET("/provinces", routes.GetProvinces)
	r.GET("/regencies/:province_id", routes.GetRegencies)
	r.GET("/district/:regencie_id", routes.GetDistrict)
	r.GET("/villages/:district_id", routes.GetVilanges)

	// Jalankan server
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
