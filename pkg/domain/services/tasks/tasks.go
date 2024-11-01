package tasks

import (
	"github.com/minhajuddin/todo/pkg/domain/models"
	taskRepo "github.com/minhajuddin/todo/pkg/domain/repos/tasks"
)

func ListTasks() []models.Task {
	return taskRepo.ListTasks()
}
