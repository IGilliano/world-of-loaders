package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"worldOfLoaders/pkg/repository/repo_models"
)

type LoaderPostgres struct {
	db *sqlx.DB
}

func NewLoader(db *sqlx.DB) *LoaderPostgres {
	return &LoaderPostgres{db: db}
}

func (l *LoaderPostgres) GetLoaderFromDB(ID int) (repo_models.Loader, error) {
	var loader repo_models.Loader
	query := fmt.Sprintf(`SELECT * FROM loaders WHERE p_id = $1`)
	err := l.db.Get(&loader, query, ID)

	return loader, err
}

func (l *LoaderPostgres) GetTasksFromDB(ID int) ([]*repo_models.Task, error) {
	var tasks []*repo_models.Task

	rows, err := l.db.Query("SELECT * FROM tasks_archive WHERE p_id = $1", ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task repo_models.Task
		err = rows.Scan(&task.Name, &task.Weight, &task.Item)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
