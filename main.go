package main

import (
	"fmt"
	"os"
	"strings"
)

var tasksFile = "tasks.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-manager-cli <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-manager-sli add <task>")
			os.Exit(1)
		}
		task := os.Args[2]
		addTask(task)
		fmt.Println("Task added succesfully. Exiting....")
		os.Exit(0)
	case "list":
	listTasks()
	default:
	fmt.Println("Unknown command. Available commands: add, list")
	}
}


func addTask(task string) {
	// read the existing tasks from the file
	existingTasks, err := os.ReadFile(tasksFile)
	if err != nil && !os.IsNotExist(err){
		fmt.Println("Error reading tasks:", err)
		os.Exit(1)
	}

	// Append new Task
	newTasks :=  append(strings.Split(string(existingTasks), "\n"), task)

	// write the updated task back to the file
	err = os.WriteFile(tasksFile, []byte(strings.Join(newTasks, "\n")), 0644)
	if err != nil {
		fmt.Println("Error writing tasks:", err)
		os.Exit(1)
	}

	fmt.Printf("Task added: %s\n", task)
}

func listTasks() {{
	// Read tasks from the file or create an empty file if it doesnt exist
	tasks, err := os.ReadFile(tasksFile)
	if err != nil && !os.IsNotExist(err) {
		// handles the error if it's not related to the file not existing 
		fmt.Println("Error reading tasks:", err)
		os.Exit(1)
	}

	// print the tasks
	if len(tasks) == 0 {
		fmt.Println("No tasks available")
	} else {
		fmt.Println("Task list:")
		fmt.Print(string(tasks))
	}
}}