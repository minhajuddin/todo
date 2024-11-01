package tasks

import "github.com/minhajuddin/todo/pkg/domain/models"

func ListTasks() []models.Task {
	return []models.Task{
		{ID: 1, Title: "Task 1"},
		{ID: 2, Title: "Task 2"},
	}
}
