package service

import (
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/repository"
)

type GetTaskByPk struct {
	repo repository.TaskRepositorier
	id   model.TaskID
	Task *model.Task
}

// コンストラクタ
func GetTaskByPkService(repo repository.TaskRepositorier, id model.TaskID) *GetTaskByPk {
	service := &GetTaskByPk{
		repo: repo,
		id:   id,
	}
	return service
}

// validate
func (gt *GetTaskByPk) Validate() error {
	return nil
}

// execute
func (gt *GetTaskByPk) Execute() error {
	task, err := gt.repo.GetByPk(gt.id)
	gt.Task = task
	return err
}
