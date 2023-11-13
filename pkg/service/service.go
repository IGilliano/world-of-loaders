package service

import (
	"worldOfLoaders/pkg/repository"
	"worldOfLoaders/pkg/repository/repo_models"
)

type Authorization interface {
	CreatePlayer(user repo_models.Player) (int, error)
	GetPlayers() ([]*repo_models.Player, error)
	CreateTasks(int) ([]int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, string, error)
}

type Loader interface {
	GetLoader(int) (repo_models.Loader, error)
	GetLoaderTasks(int) ([]*repo_models.Task, error)
}

type Client interface {
	GetClient(int) (repo_models.Client, error)
	GetClientTasks() ([]*repo_models.Task, error)
	GetTask(id int) (repo_models.Task, error)
	GetLoaders(loadersID []int) ([]*repo_models.Loader, error)
	Start(playerID int, taskID int, loadersID []int) (bool, error)
}

type Service struct {
	Authorization
	Loader
	Client
}

func NewService(rep *repository.Repository) *Service {
	return &Service{NewAuthService(rep.Authorization), NewLoaderService(rep.Loader), NewClientService(rep.Client)}
}
