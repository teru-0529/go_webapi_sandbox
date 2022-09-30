package service

import (
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/repository"
)

type ListTasks struct {
	repo  repository.TaskRepositorier
	Tasks model.Tasks
}

// コンストラクタ
func ListTasksService(repo repository.TaskRepositorier) *ListTasks {
	service := &ListTasks{
		repo: repo,
	}
	return service
}

// validate
func (lt *ListTasks) Validate() error {
	return nil
}

// execute
func (lt *ListTasks) Execute() error {
	tasks, err := lt.repo.List()
	lt.Tasks = tasks
	return err
}
