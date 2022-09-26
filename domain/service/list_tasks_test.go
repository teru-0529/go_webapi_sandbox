package service

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
)

func TestListTasks(t *testing.T) {

	now := time.Now()

	type want struct {
		resTasks model.Tasks
	}
	tests := map[string]struct {
		tasks map[model.TaskID]*model.Task
		want  want
	}{
		"ok": {
			tasks: map[model.TaskID]*model.Task{
				1: {
					ID:        1,
					Title:     "test1",
					Status:    "todo",
					CreatedAt: now,
					UpdatedAt: now,
				},
				2: {
					ID:        2,
					Title:     "test2",
					Status:    "done",
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			want: want{
				resTasks: []*model.Task{
					{
						ID:        1,
						Title:     "test1",
						Status:    "todo",
						CreatedAt: now,
						UpdatedAt: now,
					},
					{
						ID:        2,
						Title:     "test2",
						Status:    "done",
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
		},
		"empty": {
			tasks: map[model.TaskID]*model.Task{},
			want: want{
				resTasks: []*model.Task{},
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			// t.Parallel() //INFO:テストをパラレルで行うことができる

			repo := &in_memory.TaskRepository{Tasks: tt.tasks}

			service := ListTasksService(repo)

			// 実行
			_ = service.Execute()

			if diff := cmp.Diff(service.Tasks, tt.want.resTasks); diff != "" {
				t.Errorf("got differs: (-got +want)\n%s", diff)
			}
		})
	}

}
