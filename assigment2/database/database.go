package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	dt "assigment2/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func StartDB() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser, dbPassword, dbName, dbPort, dbHost :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_HOST")

	convertPort, _ := strconv.Atoi(dbPort)

	config := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", dbHost, convertPort, dbUser, dbPassword, dbName)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&dt.Items{}, &dt.Orders{})
	fmt.Println("Successfully connect to database and migrate")
}

func GetDB() *gorm.DB {
	return db
}
