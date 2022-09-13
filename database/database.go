package database

import (
	"github/com/Naithar01/go_rest_api/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:snmsung1.@tcp(localhost:3306)/rest_api?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(2)
	}

	db.AutoMigrate(&models.Category{}, &models.Post{})

	Database = db
}
