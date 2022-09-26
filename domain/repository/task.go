package repository

import "github.com/teru-0529/go_webapi_sandbox/domain/model"

type TaskRepositorier interface {
	// タスク追加
	Add(task *model.Task) (*model.Task, error)
	// 全タスク取得
	List() model.Tasks
	// タスク取得（PK指定）
	GetByPk(id model.TaskID) (*model.Task, error)
}
