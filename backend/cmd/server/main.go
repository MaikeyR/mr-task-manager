package main

import (
	"database/sql"
	"log"
	"mr-task-manager/backend/internal/config"
	"mr-task-manager/backend/internal/handlers"
	"mr-task-manager/backend/internal/repositories"
	"mr-task-manager/backend/internal/services"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Load application configuration
	appConfig := config.LoadConfig()

	// Connect to the PostgreSQL database
	database, connectionError := sql.Open("postgres", appConfig.GetDBConnectionString())
	if connectionError != nil {
		log.Fatal(connectionError)
	}

	// Create a new task repository
	taskRepository := repositories.NewTaskRepository(database)
	taskService := services.NewTaskService(taskRepository)

	// Create a new router using the Gorilla mux package
	router := mux.NewRouter()

	// Define routes and their corresponding handlers
	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	router.HandleFunc("/api/tasks", handlers.TasksHandler(taskService)).Methods("GET", "POST", "PUT", "DELETE")

	// Start the server and listen for incoming requests
	serverPort := ":8080"
	log.Println("Server started at", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, router))
}
