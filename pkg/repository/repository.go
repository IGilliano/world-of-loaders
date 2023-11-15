package repository

import (
	"github.com/jmoiron/sqlx"
	"worldOfLoaders/pkg/models"
)

type Authorization interface {
	CreatePlayer(player models.Player) (int, error)
	GetPlayer(login, password string) (models.Player, error)
	GetPlayers() ([]*models.Player, error)
}

type Loader interface {
	GetLoaderFromDB(ID int) (models.Loader, error)
	GetLoadersFromDB(loadersID []int) ([]*models.Loader, error)
	GetRegisteredLoaders() ([]*models.Loader, error)
	UpdateLoaders(loaders []*models.Loader) error
}

type Client interface {
	GetClientFromDB(ID int) (models.Client, error)
	UpdateClient(models.Client) error
}

type Task interface {
	PushTasks(tasks []models.Task) ([]int, error)
	GetTaskFromDB(id int) (models.Task, error)
	GetTasksFromDB(ID int) ([]*models.Task, error)
	GetAvailableTasksFromDB() ([]*models.Task, error)
	UpdateTasks(taskID int, loaders []int) error
}

type Repository struct {
	Authorization
	Loader
	Client
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthPostgres(db), NewLoader(db), NewClient(db), NewTask(db)}
}
