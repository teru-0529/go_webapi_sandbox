package service

import (
	"testing"
	"time"

	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/testutil"
)

func TestPostTasks(t *testing.T) {
	type want struct {
		resTask *model.Task
	}
	tests := map[string]struct {
		reqTask *model.Task
		want    want
	}{
		"ok": {
			reqTask: &model.Task{
				Title:  "title",
				Status: model.TaskStatusTodo,
				// CreatedAt:  now,
				// ModifiedAt: now,
			},
			want: want{
				resTask: &model.Task{
					ID:         1,
					Title:      "title",
					Status:     model.TaskStatusTodo,
					CreatedAt:  time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC),
					ModifiedAt: time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC),
				},
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel() //INFO:テストをパラレルで行うことができる

			service := PostTasksService(in_memory.InMemoryRepo, tt.reqTask)

			// 実行
			_ = service.Execute()
			testutil.AssertTask(t, service.Task, tt.want.resTask)
		})
	}
}
