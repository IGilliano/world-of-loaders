package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"worldOfLoaders/pkg/models"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTask(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (t *TaskPostgres) GetTaskFromDB(id int) (models.Task, error) {
	var task models.Task

	query := fmt.Sprintf(`SELECT * FROM tasks WHERE id = $1 AND available = true`)

	if err := t.db.Get(&task, query, id); err != nil {
		return task, err
	}

	return task, nil
}

func (t *TaskPostgres) GetTasksFromDB(ID int) ([]*models.Task, error) {
	var tasks []*models.Task

	rows, err := t.db.Query("SELECT t.id, t.name, t.weight, t.item FROM tasks_archive as ta JOIN tasks as t ON t.id = ta.t_id WHERE p_id = $1", ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Name, &task.Weight, &task.Item)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (t *TaskPostgres) GetAvailableTasksFromDB() ([]*models.Task, error) {
	var tasks []*models.Task

	rows, err := t.db.Query("SELECT * FROM tasks WHERE available = true")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Name, &task.Item, &task.Weight, &task.Available)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (t *TaskPostgres) PushTasks(tasks []models.Task) ([]int, error) {
	tasksID := make([]int, 0, len(tasks))
	for _, task := range tasks {
		var id int
		row := t.db.QueryRow("INSERT INTO tasks (name, item, weight, available) VALUES ($1, $2, $3, $4) RETURNING id", task.Name, task.Item, task.Weight, true)
		if err := row.Scan(&id); err != nil {
			log.Printf("Couldnt save task to database, err: %v", err.Error())
			continue
		}
		tasksID = append(tasksID, id)
	}
	return tasksID, nil
}

func (t *TaskPostgres) UpdateTasks(taskID int, loaders []int) error {
	_, err := t.db.Query("UPDATE tasks SET available = $1 WHERE id = $2", false, taskID)
	if err != nil {
		return err
	}

	for _, loader := range loaders {
		_, err = t.db.Query("INSERT INTO tasks_archive (t_id, p_id) VALUES ($1, $2)", taskID, loader)
		if err != nil {
			return err
		}
	}
	return nil
}
