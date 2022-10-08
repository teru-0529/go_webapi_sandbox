package in_memory

import (
	"time"

	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/repository"
)

// インメモリリポジトリ（シングルトン）
var InMemoryRepo = &TaskRepository{Tasks: map[model.TaskID]*model.Task{}}

type TaskRepository struct {
	// TODO:Inメモリでの検証用のためexport
	LastID model.TaskID
	Tasks  map[model.TaskID]*model.Task
}

// コンストラクタ
func NewTaskRepo() *TaskRepository {
	// シングルトンのリポジトリを返す
	return InMemoryRepo
}

// タスク追加
func (tr *TaskRepository) Add(task *model.Task) (*model.Task, error) {
	now := time.Now()
	tr.LastID++
	task.ID = tr.LastID
	task.CreatedAt = now
	task.ModifiedAt = now
	tr.Tasks[task.ID] = task
	return task, nil
}

// 全タスク取得
func (tr *TaskRepository) List() (model.Tasks, error) {
	tasks := make([]*model.Task, len(tr.Tasks))
	for i, task := range tr.Tasks {
		tasks[i-1] = task
	}
	return tasks, nil
}

// タスク取得（PK指定）
func (tr *TaskRepository) GetByPk(id model.TaskID) (*model.Task, error) {
	if task, ok := tr.Tasks[id]; ok {
		return task, nil
	}
	return nil, repository.ErrNotFound
}

// タスク修正
func (tr *TaskRepository) Patch(id model.TaskID, fields map[model.UpdateField]any) (*model.Task, error) {
	if task, ok := tr.Tasks[id]; ok && len(fields) > 0 {
		if val, ok := fields[model.F_Title]; ok {
			task.Title = val.(string)
		}
		if val, ok := fields[model.F_Status]; ok {
			task.Status = val.(model.TaskStatus)
		}
		task.ModifiedAt = time.Now()

		return task, nil
	}
	return nil, repository.ErrNotFound
}
