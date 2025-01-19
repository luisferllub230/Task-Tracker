package main

import (
	"fmt"
	"os/exec"

	"github.com/luisferllub230/task_tracker/models"
)

func executeCommand(command string) {
	cli := exec.Command("bash", "-c", command)

	if err := cli.Run(); err != nil {
		fmt.Println("\n\n Error: %s \n\n", err)
	}
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
	fmt.Println("0. Exit")
}

func printTasks(tasks []models.Task) {
	for _, task := range tasks {
		fmt.Printf("\n\n ID: %d \n Title: %s \n Description: %s \n Status: %s \n", task.Id, task.Title, task.Description, task.Status)
	}
}

func main() {
	var option int = 1
	executeCommand("clear")
	for {
		showMenuOptions()
		fmt.Scan(&option)
		switch option {
		case 1:
			executeCommand("clear")
			tasks, err := models.ListTasks()

			if err != nil {
				fmt.Println("\n\n\nError: %w", err)
				break
			}
			printTasks(tasks)
			break
		}

		if option == 0 {
			break
		}
	}

}
