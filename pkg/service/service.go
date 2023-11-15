package service

import (
	"worldOfLoaders/pkg/models"
	"worldOfLoaders/pkg/repository"
)

type Authorization interface {
	CreatePlayer(user models.Player) (int, error)
	GetPlayers() ([]*models.Player, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, string, error)
}

type Loader interface {
	GetLoader(int) (models.Loader, error)
	GetLoaders(loadersID []int) ([]*models.Loader, error)
}

type Client interface {
	GetClientInfo(int) (models.Client, []*models.Loader, error)
	Start(playerID int, params StartParams) (bool, error)
}

type Task interface {
	CreateTasks(int) ([]int, error)
	GetLoaderTasks(int) ([]*models.Task, error)
	GetClientTasks() ([]*models.Task, error)
	GetTask(id int) (models.Task, error)
}

type Service struct {
	Authorization
	Loader
	Client
	Task
}

func NewService(rep *repository.Repository) *Service {
	return &Service{NewAuthService(rep.Authorization), NewLoaderService(rep.Loader), NewClientService(rep.Client, rep.Task, rep.Loader), NewTaskService(rep)}
}
