package models

import "fmt"

type Task struct {
	Task     string `json:"task"`
	Id       int    `json:"id"`
	Priority int    `json:"priority"`
}

func (task *Task) ToString() string {
	return fmt.Sprintf("ID: %d | Task: %v | Priority: %d", task.Id, task.Task, task.Priority)
}

func PrintTasks(tasks []Task) {
	for _, task := range tasks {
		fmt.Println(task.ToString())
	}
}