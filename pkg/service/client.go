package service

import (
	"log"
	"worldOfLoaders/pkg/repository"
	"worldOfLoaders/pkg/repository/repo_models"
)

type ClientService struct {
	rep repository.Client
}

func NewClientService(rep repository.Client) *ClientService {
	return &ClientService{rep: rep}
}

func (c *ClientService) GetClient(ID int) (repo_models.Client, error) {
	return c.rep.GetClientFromDB(ID)
}

func (c *ClientService) GetClientTasks() ([]*repo_models.Task, error) {
	return c.rep.GetAvailableTasksFromDB()
}

func (c *ClientService) GetTask(id int) (repo_models.Task, error) {
	return c.rep.GetTaskFromDB(id)
}

func (c *ClientService) GetLoaders(loadersID []int) ([]*repo_models.Loader, error) {
	return c.rep.GetLoadersFromDB(loadersID)
}

func (c *ClientService) Start(playerID int, taskID int, loadersID []int) (bool, error) {
	var result bool
	var bill int
	var totalCapacity float64
	client, err := c.rep.GetClientFromDB(playerID)
	if err != nil {
		return result, err
	}
	log.Printf("Client is ready")

	task, err := c.rep.GetTaskFromDB(taskID)
	if err != nil {
		return result, err
	}
	log.Printf("Task is ready")

	loaders, err := c.rep.GetLoadersFromDB(loadersID)
	if err != nil {
		return result, err
	}
	log.Printf("Loaders is ready")

	for _, loader := range loaders {
		var lCap float64
		/*
				if loader.IsDrinking {
				lCap = float64((loader.Capacity) * ((100 - loader.Fatigue) / 100) * (loader.Fatigue / 100))
			} else {
				lCap = float64((loader.Capacity) * ((100 - loader.Fatigue) / 100))
			}
		*/
		lCap = float64((loader.Capacity) * ((100 - loader.Fatigue) / 100))
		log.Printf("Here is loader number %d with capacity:%f", loader.ID, lCap)
		bill += loader.Salary
		totalCapacity += lCap
		loader.Fatigue += 20
		if loader.IsDrinking {
			loader.Fatigue += 10
		}
	}

	log.Printf("Now, now. Player have %d dollars, bill:%d dollars", client.Fund, bill)
	if client.Fund < bill {
		client.Fund = 0
		client.InGame = false
		if err = c.rep.UpdateClient(client); err != nil {
			return result, err
		}
		return result, nil
	}

	log.Printf("Capacity of loaders: %f, but they need to take %d", totalCapacity, task.Weight)
	if totalCapacity < float64(task.Weight) {
		client.Fund = 0
		client.InGame = false
		if err = c.rep.UpdateClient(client); err != nil {
			return result, err
		}
		return result, nil
	}

	client.Fund -= bill

	if err = c.rep.UpdateClient(client); err != nil {
		return result, err
	}
	if err = c.rep.UpdateLoaders(loaders); err != nil {
		return result, err
	}

	if err = c.rep.UpdateTasks(task.ID, loadersID); err != nil {
		return result, err
	}

	return true, nil
}
