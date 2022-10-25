package database

import (
	"fmt"
	"log"

	"mygram/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init(url string) *gorm.DB {
	db, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{}, &models.Comments{}, &models.Photo{}, &models.SocialMedias{})

	fmt.Println("Database connected")

	return db
}

func GetDB() *gorm.DB {
	return db
}
