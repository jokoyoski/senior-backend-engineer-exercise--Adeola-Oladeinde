package main

import (
	"log"
	"net/http"
	"syndio-backend-app/config"
	"syndio-backend-app/database"
	"syndio-backend-app/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.InitDB()
	defer database.DB.Close()

	http.HandleFunc("/save-jobs", handlers.SaveJobsHandler)

	port := config.GetPort()
	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
