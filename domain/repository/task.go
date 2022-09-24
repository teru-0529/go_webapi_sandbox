package repository

import "github.com/teru-0529/go_webapi_sandbox/domain/model"

type TaskRepositorier interface {
	// タスク追加
	Add(task *model.Task) (model.TaskID, error)
	// 全タスク取得
	GetAll() model.Tasks
	// タスク取得（ID指定）
	GetById(id model.TaskID) (*model.Task, error)
}
