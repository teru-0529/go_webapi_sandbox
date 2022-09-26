package service

import (
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/repository"
)

type PostTasks struct {
	Repository repository.TaskRepositorier
	Task       *model.Task
}

// コンストラクタ
func PostTasksService(repositorier repository.TaskRepositorier, task *model.Task) *PostTasks {
	postTasks := &PostTasks{
		Repository: repositorier,
		Task:       task,
	}
	return postTasks
}

// validate
func (pt *PostTasks) Validate() error {
	return nil
}

// execute
func (pt *PostTasks) Execute() error {
	task, err := pt.Repository.Add(pt.Task)
	pt.Task = task
	return err
}
