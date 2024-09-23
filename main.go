package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/joshua468/team-budget-management/config"
	"github.com/joshua468/team-budget-management/handlers"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize the database
	config.InitDB()

	// Initialize router
	router := mux.NewRouter()

	// Team Routes
	router.HandleFunc("/teams", handlers.CreateTeamHandler).Methods("POST")
	router.HandleFunc("/teams", handlers.GetTeamsHandler).Methods("GET")

	// Expense Routes
	router.HandleFunc("/teams/{team_id}/expenses", handlers.AddExpenseHandler).Methods("POST")
	router.HandleFunc("/teams/{team_id}/expenses", handlers.GetExpensesByTeamHandler).Methods("GET")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
