package repositories

import (
	"database/sql"

	"crud-backend/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func (r *TaskRepository) Insert(task *models.Task) error {
	query := `INSERT INTO tasks (task) VALUES ($1);`

	_, err := r.DB.Exec(query, task.Task)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) Update(id int, task *models.Task) error {
	query := `UPDATE tasks SET task = $1, is_done = $2 WHERE id = $3;`

	_, err := r.DB.Exec(query, task.Task, task.IsDone, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) FetchAll() ([]models.Task, error) {
	query := `SELECT id, task, is_done, created_at, updated_at FROM tasks;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	tasks := []models.Task{}

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Task, &task.CreatedAt)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
