package main

import (
	"assigment2/controllers"
	"assigment2/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	database.StartDB()
	dbHandler := database.GetDB()

	controllers.RegisterRoutes(r, dbHandler)

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Assigment 2 Orders API")
	})

	r.Run(port)
}
