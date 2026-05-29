package database

import (
	"fmt"

	models "backend-umkm/models/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	const MYSQL = "root:@tcp(127.0.0.1:3305)/umkm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	DB, err = gorm.Open(mysql.Open(MYSQL), &gorm.Config{})
	if err != nil {
		panic("gagal connetc ke database: " + err.Error())

	}
	fmt.Println("Database Connected!")
	DB.AutoMigrate(&models.Category{},
		&models.Product{})
}
