package service

import (
	"errors"
	"testing"
	"time"

	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/testutil"
)

func TestGetTaskByPk(t *testing.T) {
	now := time.Now()

	tasks := map[model.TaskID]*model.Task{
		1: {
			ID:         1,
			Title:      "test1",
			Status:     "todo",
			CreatedAt:  now,
			ModifiedAt: now,
		},
		2: {
			ID:         2,
			Title:      "test2",
			Status:     "done",
			CreatedAt:  now,
			ModifiedAt: now,
		},
	}

	type want struct {
		err     error
		resTask *model.Task
	}
	tests := map[string]struct {
		tasks   map[model.TaskID]*model.Task
		modelId model.TaskID
		want    want
	}{
		"ok": {
			tasks:   tasks,
			modelId: model.TaskID(1),
			want: want{
				err: nil,
				resTask: &model.Task{
					ID:         1,
					Title:      "test1",
					Status:     model.TaskStatusTodo,
					CreatedAt:  now,
					ModifiedAt: now,
				},
			},
		},
		"not_found": {
			tasks:   tasks,
			modelId: model.TaskID(10),
			want: want{
				err:     in_memory.ErrNotFound,
				resTask: nil,
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel() //INFO:テストをパラレルで行うことができる

			repo := &in_memory.TaskRepository{Tasks: tt.tasks}
			service := GetTaskByPkService(repo, tt.modelId)

			// 実行
			err := service.Execute()

			testutil.AssertTask(t, tt.want.resTask, service.Task)
			if !errors.Is(tt.want.err, err) {
				t.Errorf("%s error\nwant: %+v\ngot : %+v\n", t.Name(), tt.want.err, err)
			}
		})
	}
}
