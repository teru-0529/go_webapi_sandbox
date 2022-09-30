package repository

import "github.com/teru-0529/go_webapi_sandbox/domain/model"

type TaskRepositorier interface {
	// タスク追加
	Add(task *model.Task) (*model.Task, error)
	// 全タスク取得
	List() (model.Tasks, error)
	// タスク取得（PK指定）
	GetByPk(id model.TaskID) (*model.Task, error)
	// タスク修正
	Patch(id model.TaskID, fields map[model.UpdateField]any) (*model.Task, error)
}
