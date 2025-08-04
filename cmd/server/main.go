package main

import (
	"log"
	"owwi/pkg/api"
	databases "owwi/pkg/database"

	"github.com/joho/godotenv"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found, proceeding without it")
	}

	if _, err := databases.SetupMongoDB(); err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	router := api.NewRouter()

	router.Run("localhost:8081")
}
