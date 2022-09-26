package service

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
)

func TestPostTasks(t *testing.T) {
	now := time.Now()

	type want struct {
		resTask *model.Task
	}
	tests := map[string]struct {
		reqTask *model.Task
		want    want
	}{
		"ok": {
			reqTask: &model.Task{
				Title:     "title",
				Status:    model.TaskStatusTodo,
				CreatedAt: now,
				UpdatedAt: now,
			},
			want: want{
				resTask: &model.Task{
					ID:        1,
					Title:     "title",
					Status:    model.TaskStatusTodo,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			// t.Parallel() //INFO:テストをパラレルで行うことができる

			task := tt.reqTask
			service := PostTasksService(in_memory.InMemoryRepo, task)

			// 実行
			_ = service.Execute()

			if diff := cmp.Diff(service.Task, tt.want.resTask); diff != "" {
				t.Errorf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}
