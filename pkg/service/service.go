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

type Service struct {
	Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{NewAuthService(rep.Authorization)}
}
