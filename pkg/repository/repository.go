package repository

import (
	"github.com/jmoiron/sqlx"
	"worldOfLoaders/pkg/repository/repo_models"
)

type Authorization interface {
	CreatePlayer(player repo_models.Player) (int, error)
	GetPlayers() ([]*repo_models.Player, error)
	GetPlayer(login, password string) (repo_models.Player, error)
	PushTasks(tasks []repo_models.Task) ([]int, error)
}

type Loader interface {
	GetLoaderFromDB(ID int) (repo_models.Loader, error)
	GetTasksFromDB(ID int) ([]*repo_models.Task, error)
}

type Client interface {
	GetClientFromDB(ID int) (repo_models.Client, error)
	GetAvailableTasksFromDB() ([]*repo_models.Task, error)
	GetTaskFromDB(id int) (repo_models.Task, error)
	GetLoadersFromDB(loadersID []int) ([]*repo_models.Loader, error)
	UpdateClient(repo_models.Client) error
	UpdateLoaders(loaders []*repo_models.Loader) error
	UpdateTasks(taskID int, loaders []int) error
}

type Repository struct {
	Authorization
	Loader
	Client
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthPostgres(db), NewLoader(db), NewClient(db)}
}
