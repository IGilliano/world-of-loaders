package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"worldOfLoaders/pkg/repository/repo_models"
)

type ClientPostgres struct {
	db *sqlx.DB
}

func NewClient(db *sqlx.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (c *ClientPostgres) GetClientFromDB(ID int) (repo_models.Client, error) {
	var client repo_models.Client
	query := fmt.Sprintf(`SELECT * FROM clients WHERE p_id = $1`)
	err := c.db.Get(&client, query, ID)

	return client, err
}

func (c *ClientPostgres) GetAvailableTasksFromDB() ([]*repo_models.Task, error) {
	var tasks []*repo_models.Task

	rows, err := c.db.Query("SELECT * FROM tasks WHERE available = true")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task repo_models.Task
		err = rows.Scan(&task.ID, &task.Name, &task.Item, &task.Weight, &task.Available)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (c *ClientPostgres) GetTaskFromDB(id int) (repo_models.Task, error) {
	var task repo_models.Task

	query := fmt.Sprintf(`SELECT * FROM tasks WHERE id = $1 AND available = true`)

	if err := c.db.Get(&task, query, id); err != nil {
		return task, err
	}

	return task, nil
}

func (c *ClientPostgres) UpdateClient(client repo_models.Client) error {
	_, err := c.db.Query("UPDATE clients SET fund = $1, in_game = $2 WHERE p_id = $3", client.Fund, client.InGame, client.ID)
	return err
}

func (c *ClientPostgres) GetLoadersFromDB(loadersID []int) ([]*repo_models.Loader, error) {
	loaders := make([]*repo_models.Loader, 0)
	for _, id := range loadersID {
		var loader repo_models.Loader
		query := fmt.Sprintf(`SELECT * FROM loaders WHERE p_id = $1`)
		err := c.db.Get(&loader, query, id)
		if err != nil {
			return nil, err
		}
		loaders = append(loaders, &loader)
	}
	return loaders, nil
}

func (c *ClientPostgres) UpdateLoaders(loaders []*repo_models.Loader) error {
	for _, loader := range loaders {
		_, err := c.db.Query("UPDATE loaders SET fatigue = $1 WHERE p_id = $2", loader.Fatigue, loader.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ClientPostgres) UpdateTasks(taskID int, loaders []int) error {
	_, err := c.db.Query("UPDATE tasks SET available = $1 WHERE id = $2", false, taskID)
	if err != nil {
		return err
	}

	for _, loader := range loaders {
		_, err := c.db.Query("INSERT INTO tasks_archive (t_id, p_id) VALUES ($1, $2)", taskID, loader)
		if err != nil {
			return err
		}
	}
	return nil
}
