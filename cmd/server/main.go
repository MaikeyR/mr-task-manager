package main

import (
	"database/sql"
	"log"
	"mr-task-manager/internal/config"
	"mr-task-manager/internal/handlers"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
    cfg := config.LoadConfig()
    db, err := sql.Open("postgres", cfg.GetDBConnectionString())
    if err != nil {
        log.Fatal(err)
    }

    router := mux.NewRouter()
    router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
    router.HandleFunc("/tasks", handlers.TasksHandler(db)).Methods("GET", "POST", "PUT", "DELETE")

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
