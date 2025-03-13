package main

import (
	"backend/database"
	"backend/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.ConnectDB()

	r := routes.SetupRouter(database.DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(r.Run(":" + port))
}
