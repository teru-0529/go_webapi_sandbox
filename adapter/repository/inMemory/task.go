package inmemory

import (
	"errors"

	"github.com/teru-0529/go_webapi_sandbox/domain/model"
)

var (
	Tasks       = &TaskRepository{Tasks: map[model.TaskID]*model.Task{}}
	ErrNotFound = errors.New("not found")
)

type TaskRepository struct {
	// TODO:DB無しの検証用のためexport
	LastID model.TaskID
	Tasks  map[model.TaskID]*model.Task
}

// タスク追加
func (tr *TaskRepository) Add(task *model.Task) (model.TaskID, error) {
	tr.LastID++
	task.ID = tr.LastID
	tr.Tasks[task.ID] = task
	return task.ID, nil
}

// 全タスク取得
func (tr *TaskRepository) GetAll() model.Tasks {
	tasks := make([]*model.Task, len(tr.Tasks))
	for i, t := range tr.Tasks {
		tasks[i-1] = t
	}
	return tasks
}

// タスク取得（ID指定）
func (tr *TaskRepository) GetById(id model.TaskID) (*model.Task, error) {
	if task, ok := tr.Tasks[id]; ok {
		return task, nil
	}
	return nil, ErrNotFound
}
