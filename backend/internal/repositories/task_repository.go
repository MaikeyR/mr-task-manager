package repositories

import (
	"database/sql"
	"mr-task-manager/backend/internal/models"
	"time"
)

type TaskRepository struct {
	database *sql.DB
}

func NewTaskRepository(database *sql.DB) *TaskRepository {
	return &TaskRepository{database: database}
}

func (repository *TaskRepository) CreateTask(task *models.Task) error {
	query := `INSERT INTO tasks (name, completed, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	now := time.Now()
	err := repository.database.QueryRow(query, task.Name, task.Completed, now, now).Scan(&task.ID)
	if err != nil {
		return err
	}
	task.CreatedAt = now
	task.UpdatedAt = now
	return nil
}

func (repository *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	query := `SELECT id, name, completed, created_at, updated_at FROM tasks WHERE id = $1`
	row := repository.database.QueryRow(query, id)

	task := &models.Task{}
	err := row.Scan(&task.ID, &task.Name, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return task, nil
}

func (repository *TaskRepository) GetAllTasks() ([]*models.Task, error) {
	query := `SELECT id, name, completed, created_at, updated_at FROM tasks`
	rows, err := repository.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(&task.ID, &task.Name, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (repository *TaskRepository) UpdateTask(task *models.Task) error {
	query := `UPDATE tasks SET name = $1, completed = $2, updated_at = $3 WHERE id = $4`
	now := time.Now()
	_, err := repository.database.Exec(query, task.Name, task.Completed, now, task.ID)
	if err != nil {
		return err
	}
	task.UpdatedAt = now
	return nil
}

func (repository *TaskRepository) DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := repository.database.Exec(query, id)
	return err
}
