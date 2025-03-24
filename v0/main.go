package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // a slice to hold my tasks
	tasks := []string{}
	scanner := bufio.NewScanner(os.Stdin) 

	for {
        fmt.Println("What do you want to do? (add, list, quit)")
        scanner.Scan()       
        command := scanner.Text() 

        if command == "add" {
            fmt.Println("Enter a task:")
            scanner.Scan()
            task := scanner.Text()
            tasks = append(tasks, task)
            fmt.Println("Task added!")
        } else if command == "list" {
            if len(tasks) == 0 {
                fmt.Println("No tasks yet!")
            }
            for i, task := range tasks {
                fmt.Println(i+1, task)
            }
        } else if command == "quit" {
            fmt.Println("Bye!")
            break
        } else {
            fmt.Println("I don’t understand that command.")
        }
    }


}