// Package Task handles the Task struct and it's methods
package models

import "fmt"

type Task struct {
	Task string `json:"task"`
	Id   int    `json:"id"`
}

// ToString formats a task to be printed
func (task *Task) ToString() string {
	return fmt.Sprintf("ID: %d | Task: %v", task.Id, task.Task)
}

// PrintTasks prints the formatted tasks
func PrintTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks registered yet!")
		return
	}

	for _, task := range tasks {
		fmt.Println(task.ToString())
	}
}
