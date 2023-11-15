package service

import (
	"log"
	"worldOfLoaders/pkg/models"
	"worldOfLoaders/pkg/repository"
)

type StartParams struct {
	Task    int   `json:"task"`
	Loaders []int `json:"loaders"`
}

type ClientService struct {
	repClient  repository.Client
	repTask    repository.Task
	repLoaders repository.Loader
}

func NewClientService(repC repository.Client, repT repository.Task, repL repository.Loader) *ClientService {
	return &ClientService{repClient: repC, repTask: repT, repLoaders: repL}
}

func (c *ClientService) GetClientInfo(ID int) (models.Client, []*models.Loader, error) {
	clients, err := c.repClient.GetClientFromDB(ID)
	if err != nil {
		return models.Client{}, []*models.Loader{}, err
	}
	loaders, err := c.repLoaders.GetRegisteredLoaders()
	if err != nil {
		return models.Client{}, []*models.Loader{}, err
	}
	return clients, loaders, nil
}

func (c *ClientService) Start(playerID int, params StartParams) (bool, error) {
	var bill int
	var totalCapacity float64

	client, err := c.repClient.GetClientFromDB(playerID)
	if err != nil {
		return false, err
	}
	log.Printf("Client is ready")

	task, err := c.repTask.GetTaskFromDB(params.Task)
	if err != nil {
		return false, err
	}
	log.Printf("Task is ready")

	loaders, err := c.repLoaders.GetLoadersFromDB(params.Loaders)
	if err != nil {
		return false, err
	}
	log.Printf("Loaders is ready")

	for _, loader := range loaders {
		var lCap float64
		lCap = float64(loader.Capacity) * float64(100-loader.Fatigue) / 100
		log.Printf("Here is loader number %d with capacity:%f", loader.ID, lCap)
		bill += loader.Salary
		totalCapacity += lCap
		loader.Fatigue += 20
		if loader.IsDrinking {
			loader.Fatigue += 10
		}
		if loader.Fatigue > 100 {
			loader.Fatigue = 100
		}
	}

	log.Printf("Now, now. Player have %d dollars, bill:%d dollars", client.Fund, bill)
	log.Printf("Capacity of loaders: %f. They need to lift %d", totalCapacity, task.Weight)
	if client.Fund < bill || totalCapacity < float64(task.Weight) {
		client.Fund = 0
		client.InGame = false
		if err = c.repClient.UpdateClient(client); err != nil {
			return false, err
		}
		return false, nil
	}

	client.Fund -= bill

	if err = c.repClient.UpdateClient(client); err != nil {
		return false, err
	}
	if err = c.repLoaders.UpdateLoaders(loaders); err != nil {
		return false, err
	}
	if err = c.repTask.UpdateTasks(task.ID, params.Loaders); err != nil {
		return false, err
	}
	return true, nil
}
