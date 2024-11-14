package main

import (
	"log"
	"os"

	"github.com/Norgate-AV/simple-logging-service/db"
	"github.com/Norgate-AV/simple-logging-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	// DefaultPort is the default port for the server
	DefaultPort = "3000"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	if err := db.Initialize(); err != nil {
		log.Printf("Error initializing database: %v\n", err)
		os.Exit(1)
	}

	server := gin.Default()
	routes.Register(server)

	if err := server.Run(":" + port); err != nil {
		log.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
