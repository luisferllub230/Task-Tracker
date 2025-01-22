package main

import (
	"fmt"
	"reflect"

	"github.com/luisferllub230/task_tracker/models"
)

// TODO: DONT WORK
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func showMenuOptions() {
	fmt.Println("\n\nHello what would you like to do?")
	fmt.Println("1. List all tasks")
	fmt.Println("2. List all tasks that are done")
	fmt.Println("3. List all tasks that are not done")
	fmt.Println("4. List all tasks are in progress")
	fmt.Println("5. Mark task as progress or done")
	fmt.Println("-------------------------------")
	fmt.Println("6. Create a new task")
	fmt.Println("7. Update a task")
	fmt.Println("8. Delete a task")
	fmt.Println("-------------------------------")
	fmt.Println("9. Find a task by id")
	fmt.Println("0. Exit")
}

func printTasks(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("\n\nNo tasks to show\n\n")
		return
	}
	for _, task := range tasks {
		fmt.Printf("\n\n ID: %d \n Title: %s \n Description: %s \n Status: %s \n", task.Id, task.Title, task.Description, task.Status)
	}
}

func main() {
	var option int = 0
	clearScreen()
	for {
		showMenuOptions()
		fmt.Scan(&option)
		switch option {
		case 1:
			clearScreen()
			tasks, err := models.ListTasks()

			if err != nil {
				fmt.Println("\n\n\nError: %w", err)
				break
			}
			printTasks(tasks)
			break

		case 2:
			clearScreen()
			tasks, err := models.ListTasksByInStatus([]any{"done"})

			if err != nil {
				fmt.Println("\n\n\nError: %w", err)
				break
			}
			printTasks(tasks)
			break

		case 3:
			clearScreen()
			tasks, err := models.ListTasksByInStatus([]any{"progress", "new"})

			if err != nil {
				fmt.Println("\n\n\nError: %w", err)
				break
			}
			printTasks(tasks)
			break

		case 4:
			clearScreen()
			tasks, err := models.ListTasksByInStatus([]any{"progress"})

			if err != nil {
				fmt.Println("\n\n\nError: %w", err)
				break
			}
			printTasks(tasks)
			break

		case 7:
			clearScreen()
			var id int = 0
			fmt.Println("\n\nInsert the id of the task do you want to update. Press 0 to show all tasks")
			fmt.Scan(&id)

			if id == 0 {
				clearScreen()
				tasks, err := models.ListTasks()

				if err != nil {
					fmt.Println("\n\n\nError: ", err)
				}
				printTasks(tasks)
				break
			}

			task, err := models.FindTaskById(id)
			if err != nil {
				fmt.Println("\n\n\nError: %w", err.Error())
				break
			}

			reflectTask := reflect.ValueOf(&task).Elem()
			numberOfFields := reflectTask.NumField()
			i := 1
			for {
				field := reflectTask.Field(i)
				fieldName := reflectTask.Type().Field(i).Name
				fmt.Printf("Current value of the field %s: %v\n", fieldName, field.Interface())
				fmt.Printf("New value: ")
				var newValue string = ""
				fmt.Scan(&newValue)
				reflectTask.Field(i).Set(reflect.ValueOf(newValue))
				i++
				if i >= numberOfFields {
					break
				}
			}

			updateTask, _ := models.UpdateTask(task)
			printTasks([]models.Task{updateTask})
			break
		case 9:
			clearScreen()
			var id int = 0
			fmt.Println("\n\nInsert the id of the task you want to find else press 0 to list all tasks")
			fmt.Scan(&id)

			if id == 0 {
				clearScreen()
				tasks, err := models.ListTasks()

				if err != nil {
					fmt.Println("\n\n\nError: %w", err)
					break
				}

				printTasks(tasks)
				break
			}

			tasks, err := models.FindTaskById(id)

			if err != nil {
				fmt.Println("\n\n\nError: %w", err)
				break
			}

			if tasks.Id == 0 {
				fmt.Println("\n\n\nTask not found")
				break
			}

			printTasks([]models.Task{tasks})
			break
		}

		if option == 0 {
			break
		}
	}

}
