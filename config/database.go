package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:rafel_root@tcp(127.0.0.1:3307)/db_wilayah_indonesia?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Mematikan pluralisasi otomatis
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	DB = database
}
