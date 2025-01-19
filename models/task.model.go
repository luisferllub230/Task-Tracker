package models

import (
	"github.com/luisferllub230/task_tracker/services"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func CreateTask(task Task) (Task, error) {
	var newTask, err = services.Create(task)
	if err != nil {
		return Task{}, err
	}
	return newTask.(Task), nil
}

func ListTasks() ([]Task, error) {
	var tasks, err = services.List(Task{})

	if err != nil {
		return []Task{}, err
	}

	var taskList []Task
	for _, task := range tasks {
		taskMap := task.(map[string]interface{})
		taskList = append(taskList, Task{
			Id:          int(taskMap["id"].(float64)),
			Title:       taskMap["title"].(string),
			Description: taskMap["description"].(string),
			Status:      taskMap["status"].(string),
		})
	}

	return taskList, nil
}
