package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/testutil"
)

func TestListTasks(t *testing.T) {
	type want struct {
		status  int
		resFile string
	}
	tests := map[string]struct {
		tasks map[model.TaskID]*model.Task
		want  want
	}{
		"ok": {
			tasks: map[model.TaskID]*model.Task{
				1: {
					ID:         1,
					Title:      "test1",
					Status:     "todo",
					CreatedAt:  time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC),
					ModifiedAt: time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC),
				},
				2: {
					ID:         2,
					Title:      "test2",
					Status:     "done",
					CreatedAt:  time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC),
					ModifiedAt: time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC),
				},
			},
			want: want{
				status:  http.StatusOK,
				resFile: "testdata/list_tasks/ok_res.json.golden",
			},
		},
		"empty": {
			tasks: map[model.TaskID]*model.Task{},
			want: want{
				status:  http.StatusOK,
				resFile: "testdata/list_tasks/empty_res.json.golden",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel() //INFO:テストをパラレルで行うことができる

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/tasks", nil)

			sut := ListTasks{Repository: &in_memory.TaskRepository{Tasks: tt.tasks}}
			sut.ServeHTTP(w, r)

			res := w.Result()
			testutil.AssertResponse(t, res, tt.want.status, testutil.LoadFile(t, tt.want.resFile))
		})
	}
}
