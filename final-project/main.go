package main

import (
	"log"
	"mygram/app/routers"
	db "mygram/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUrl := os.Getenv("POSTGRES_URL")
	port := os.Getenv("PORT")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	db.Init(dbUrl)

	routers.SetupRoutes(r)

	r.Run(port)
}
