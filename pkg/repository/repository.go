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
}

type Client interface {
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthPostgres(db)}
}
