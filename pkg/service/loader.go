package service

import (
	"worldOfLoaders/pkg/models"
	"worldOfLoaders/pkg/repository"
)

type LoaderService struct {
	rep repository.Loader
}

func NewLoaderService(rep repository.Loader) *LoaderService {
	return &LoaderService{rep: rep}
}

func (l *LoaderService) GetLoader(ID int) (models.Loader, error) {
	return l.rep.GetLoaderFromDB(ID)
}

func (l *LoaderService) GetLoaders(loadersID []int) ([]*models.Loader, error) {
	return l.rep.GetLoadersFromDB(loadersID)
}
