package service

import (
	"worldOfLoaders/pkg/repository"
	"worldOfLoaders/pkg/repository/repo_models"
)

type LoaderService struct {
	rep repository.Loader
}

func NewLoaderService(rep repository.Loader) *LoaderService {
	return &LoaderService{rep: rep}
}

func (l *LoaderService) GetLoader(ID int) (repo_models.Loader, error) {
	return l.rep.GetLoaderFromDB(ID)
}

func (l *LoaderService) GetLoaderTasks(ID int) ([]*repo_models.Task, error) {
	return l.rep.GetTasksFromDB(ID)
}
