package main

import (
	"fmt"
	"os"
	"strings"
)
// Task represents a single todo item
type Task struct {
	ID        int
	Desc      string
	Completed bool
}

// TodoList holds all tasks
type TodoList struct {
	Tasks []Task
}

func (tl *TodoList) addTask(desc string) {
	// Generating a new ID based on the current length of Tasks
	id := len(tl.Tasks) + 1
	task := Task{ID: id, Desc: desc, Completed: false}
	tl.Tasks = append(tl.Tasks, task)
	fmt.Printf("Added task: %s (ID: %d)\n", desc, id)
}

func (tl *TodoList) listTasks() {
	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tl.Tasks {
		status := " "
		if task.Completed {
			status = "✓"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Desc)
	}
}

func main() {
	todo := &TodoList{}

	// Simple CLI loop
	for {
		fmt.Println("\nCommands: add <task>, list, quit")
		var input string
		fmt.Scanln(&input)
		parts := strings.SplitN(input, " ", 2)

		switch parts[0] {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Please provide a task description.")
				continue
			}
			todo.addTask(parts[1])
		case "list":
			todo.listTasks()
		case "quit":
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Unknown command.")
		}
	}
}