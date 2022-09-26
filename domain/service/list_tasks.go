package service

import (
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/repository"
)

type ListTasks struct {
	Repository repository.TaskRepositorier
	Tasks      model.Tasks
}

// コンストラクタ
func ListTasksService(repositorier repository.TaskRepositorier) *ListTasks {
	listTasks := &ListTasks{
		Repository: repositorier,
	}
	return listTasks
}

// validate
func (lt *ListTasks) Validate() error {
	return nil
}

// execute
func (lt *ListTasks) Execute() error {
	tasks, err := lt.Repository.List()
	lt.Tasks = tasks
	return err
}
