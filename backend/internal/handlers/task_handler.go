package handlers

import (
	"encoding/json"
	"mr-task-manager/backend/internal/models"
	"mr-task-manager/backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func TasksHandler(service *services.TaskService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			tasks, err := service.GetAllTasks()
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(writer).Encode(tasks)
		case http.MethodPost:
			var task models.Task
			if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			if err := service.CreateTask(&task); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			writer.WriteHeader(http.StatusCreated)
			json.NewEncoder(writer).Encode(task)
		case http.MethodPut:
			var task models.Task
			if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			if err := service.UpdateTask(&task); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(task)
		case http.MethodDelete:
			vars := mux.Vars(request)
			id, err := strconv.Atoi(vars["id"])
			if err != nil {
				http.Error(writer, "Invalid task ID", http.StatusBadRequest)
				return
			}
			if err := service.DeleteTask(id); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			writer.WriteHeader(http.StatusNoContent)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func IndexHandler(writer http.ResponseWriter, r *http.Request) {
	// Placeholder function for IndexHandler
	writer.Write([]byte("Index Handler Placeholder"))
}
