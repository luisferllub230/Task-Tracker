package models

import (
	"errors"

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

func ListTasksByInStatus(status []any) ([]Task, error) {
	tasks, err := ListTasks()

	if err != nil {
		return []Task{}, err
	}

	var taskList []Task
	for _, task := range tasks {
		if !isInList(task.Status, status) {
			continue
		}

		taskList = append(taskList, task)
	}

	return taskList, nil
}

func FindTaskById(id int) (Task, error) {
	var task, err = services.Read(Task{Id: id})
	if err != nil {
		return Task{}, err
	}

	if task == nil {
		return Task{}, errors.New("We could not find the task")
	}

	var taskMap = task.(map[string]interface{})
	findTask := Task{
		Id:          int(taskMap["id"].(float64)),
		Title:       taskMap["title"].(string),
		Description: taskMap["description"].(string),
		Status:      taskMap["status"].(string),
	}
	return findTask, nil
}

func UpdateTask(task Task) (Task, error) {
	var updatedTask, err = services.Update(task)
	if err != nil {
		return Task{}, err
	}
	return updatedTask.(Task), nil
}

func isInList(value any, list []any) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
