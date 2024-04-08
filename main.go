package main

import (
	"dizeto-backend/app/router"
	"dizeto-backend/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	port := ":" + os.Getenv("PORT")

	// Initialize database connection and perform auto migrate
	db, err := config.InitDB()
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	// Initialize Gin router
	r := gin.Default()

	// Setup router
	router.SetupRouter(r, db)

	// Run the application
	r.Run(port)
}
