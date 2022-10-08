package service

import (
	"errors"
	"testing"
	"time"

	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/repository"
	"github.com/teru-0529/go_webapi_sandbox/testutil"
)

func TestPatchTaskForStatus(t *testing.T) {
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
		status  model.TaskStatus
		want    want
	}{
		"ok": {
			tasks:   tasks,
			modelId: model.TaskID(1),
			status:  model.TaskStatusDone,
			want: want{
				err: nil,
				resTask: &model.Task{
					ID:         1,
					Title:      "test1",
					Status:     model.TaskStatusDone,
					CreatedAt:  now,
					ModifiedAt: now,
				},
			},
		},
		"not_found": {
			tasks:   tasks,
			modelId: model.TaskID(10),
			status:  model.TaskStatusDone,
			want: want{
				err:     repository.ErrNotFound,
				resTask: nil,
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel() //INFO:テストをパラレルで行うことができる

			repo := &in_memory.TaskRepository{Tasks: tt.tasks}
			service := PatchTaskForStatusService(repo, tt.modelId, tt.status)

			// 実行
			err := service.Execute()

			testutil.AssertTask(t, tt.want.resTask, service.Task)
			if !errors.Is(tt.want.err, err) {
				t.Errorf("%s error\nwant: %+v\ngot : %+v\n", t.Name(), tt.want.err, err)
			}
		})
	}
}
