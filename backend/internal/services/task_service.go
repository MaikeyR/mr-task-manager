package services

import (
	"mr-task-manager/backend/internal/models"
	"mr-task-manager/backend/internal/repositories"
)

type TaskService struct {
	repository *repositories.TaskRepository
}

func NewTaskService(repository *repositories.TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (service *TaskService) CreateTask(task *models.Task) error {
	return service.repository.CreateTask(task)
}

func (service *TaskService) GetTaskByID(id int) (*models.Task, error) {
	return service.repository.GetTaskByID(id)
}

func (service *TaskService) GetAllTasks() ([]*models.Task, error) {
	return service.repository.GetAllTasks()
}

func (service *TaskService) UpdateTask(task *models.Task) error {
	return service.repository.UpdateTask(task)
}

func (service *TaskService) DeleteTask(id int) error {
	return service.repository.DeleteTask(id)
}
