package service

import (
	"math/rand"
	"worldOfLoaders/pkg/models"
	"worldOfLoaders/pkg/repository"
)

type TaskService struct {
	rep repository.Task
}

func NewTaskService(rep repository.Task) *TaskService {
	return &TaskService{rep: rep}
}

func (t *TaskService) CreateTasks(n int) ([]int, error) {
	tasks := make([]models.Task, 0)
	for i := 0; i < n; i++ {
		nameInt := rand.Int31n(3)
		name := models.TaskNames[nameInt]
		itemsInt := rand.Int31n(4)
		item := models.ItemNames[itemsInt]
		weight := rand.Int31n(80-10) + 10
		task := models.Task{Name: name, Item: item, Weight: int(weight)}
		tasks = append(tasks, task)
	}

	return t.rep.PushTasks(tasks)
}

func (t *TaskService) GetLoaderTasks(ID int) ([]*models.Task, error) {
	return t.rep.GetTasksFromDB(ID)
}

func (t *TaskService) GetClientTasks() ([]*models.Task, error) {
	return t.rep.GetAvailableTasksFromDB()
}

func (t *TaskService) GetTask(id int) (models.Task, error) {
	return t.rep.GetTaskFromDB(id)
}
