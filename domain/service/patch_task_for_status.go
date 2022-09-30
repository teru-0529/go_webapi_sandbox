package service

import (
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/repository"
)

type PatchTaskForStatus struct {
	repo   repository.TaskRepositorier
	id     model.TaskID
	status model.TaskStatus
	Task   *model.Task
}

// コンストラクタ
func PatchTaskForStatusService(repo repository.TaskRepositorier, id model.TaskID, status model.TaskStatus) *PatchTaskForStatus {
	service := &PatchTaskForStatus{
		repo:   repo,
		id:     id,
		status: status,
	}
	return service
}

// validate
func (pts *PatchTaskForStatus) Validate() error {
	return nil
}

// execute
func (pts *PatchTaskForStatus) Execute() error {
	m := map[model.UpdateField]any{
		model.F_Status: pts.status,
	}
	task, err := pts.repo.Patch(pts.id, m)
	pts.Task = task
	return err
}
