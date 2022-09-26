package in_memory

import (
	"errors"

	"github.com/teru-0529/go_webapi_sandbox/domain/model"
)

var (
	// インメモリリポジトリ（シングルトン）
	InMemoryRepo = &TaskRepository{Tasks: map[model.TaskID]*model.Task{}}
	// データなしエラー
	ErrNotFound = errors.New("not found")
)

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
	tr.LastID++
	task.ID = tr.LastID
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
	return nil, ErrNotFound
}
