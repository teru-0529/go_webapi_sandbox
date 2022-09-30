package service

import (
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/repository"
)

type PostTasks struct {
	repo repository.TaskRepositorier
	Task *model.Task
}

// コンストラクタ
func PostTasksService(repo repository.TaskRepositorier, task *model.Task) *PostTasks {
	service := &PostTasks{
		repo: repo,
		Task: task,
	}
	return service
}

// validate
func (pt *PostTasks) Validate() error {
	return nil
}

// execute
func (pt *PostTasks) Execute() error {
	task, err := pt.repo.Add(pt.Task)
	pt.Task = task
	return err
}
