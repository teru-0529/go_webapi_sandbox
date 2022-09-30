package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/testutil"
)

func TestGetTaskByPk(t *testing.T) {
	t.Parallel()

	tasks := map[model.TaskID]*model.Task{
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
	}

	type want struct {
		status  int
		resFile string
	}
	tests := map[string]struct {
		tasks       map[model.TaskID]*model.Task
		pathParamId string
		want        want
	}{
		"ok": {
			tasks:       tasks,
			pathParamId: "1",
			want: want{
				status:  http.StatusOK,
				resFile: "testdata/get_task_by_pk/ok_res.json.golden",
			},
		},
		"badRequest": {
			tasks:       tasks,
			pathParamId: "ABC",
			want: want{
				status:  http.StatusBadRequest,
				resFile: "testdata/get_task_by_pk/bad_res.json.golden",
			},
		},
		"NotFound": {
			tasks:       tasks,
			pathParamId: "100",
			want: want{
				status:  http.StatusNotFound,
				resFile: "testdata/get_task_by_pk/not_found_res.json.golden",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel() //INFO:テストをパラレルで行うことができる

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodGet,
				"/tasks/{id}",
				nil,
			)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.pathParamId)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			sut := GetTaskByPk{
				Repository: &in_memory.TaskRepository{Tasks: tt.tasks},
				Validator:  validator.New(),
			}
			sut.ServeHTTP(w, r)

			res := w.Result()
			testutil.AssertResponse(t, res, tt.want.status, testutil.LoadFile(t, tt.want.resFile))
		})
	}
}
