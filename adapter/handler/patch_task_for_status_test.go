package handler

import (
	"bytes"
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

func TestPatchTaskForStatus(t *testing.T) {
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
		reqFile     string
		want        want
	}{
		"ok": {
			tasks:       tasks,
			pathParamId: "1",
			reqFile:     "testdata/patch_task_for_status/ok_res.json.golden",
			want: want{
				status:  http.StatusOK,
				resFile: "testdata/patch_task_for_status/ok_res.json.golden",
			},
		},
		"badRequest": {
			tasks:       tasks,
			pathParamId: "ABC",
			reqFile:     "testdata/patch_task_for_status/bad_req.json.golden",
			want: want{
				status:  http.StatusBadRequest,
				resFile: "testdata/patch_task_for_status/bad_res.json.golden",
			},
		},
		"NotFound": {
			tasks:       tasks,
			pathParamId: "100",
			reqFile:     "testdata/patch_task_for_status/not_found_req.json.golden",
			want: want{
				status:  http.StatusNotFound,
				resFile: "testdata/patch_task_for_status/not_found_res.json.golden",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel() //INFO:テストをパラレルで行うことができる

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPatch,
				"/tasks/{id}",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.pathParamId)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			sut := PatchTaskForStatus{
				Repository: &in_memory.TaskRepository{Tasks: tt.tasks},
				Validator:  validator.New(),
			}
			sut.ServeHTTP(w, r)

			res := w.Result()
			testutil.AssertResponse(t, res, tt.want.status, testutil.LoadFile(t, tt.want.resFile))
		})
	}
}
