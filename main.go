package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)
// Task represents a single todo item
type Task struct {
	ID        int   `json:"id"`
	Desc      string `json:"desc"`
	Completed bool  `json:"completed"`
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
	tl.saveToFile("tasks.json")
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

func (tl *TodoList) saveToFile(filename string) error {
	data, err := json.MarshalIndent(tl.Tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func loadFromFile(filename string) (*TodoList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &TodoList{}, nil // Returning an empty list if file doesn’t exist
		}
		return nil, err
	}
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return &TodoList{Tasks: tasks}, nil
}

func main() {

	todo, err := loadFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

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