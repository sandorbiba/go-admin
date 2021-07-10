package database

import (
	"github.com/sandorbiba/go-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:Sanyi1111@/go_admin"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	DB = database

	database.AutoMigrate(&models.User{})
}