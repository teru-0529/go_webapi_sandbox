package hander

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/testutil"
)

func TestPostTasks(t *testing.T) {
	t.Parallel()

	type want struct {
		status  int
		resFile string
	}
	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/post_tasks/ok_req.json.golden",
			want: want{
				status:  http.StatusCreated,
				resFile: "testdata/post_tasks/ok_res.json.golden",
			},
		},
		"badRequest": {
			reqFile: "testdata/post_tasks/bad_req.json.golden",
			want: want{
				status:  http.StatusBadRequest,
				resFile: "testdata/post_tasks/bad_res.json.golden",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			// t.Parallel() INFO:テストをパラレルで行うことができる

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)

			// log.Print(100) //INFO:
			// log.Print(r)   //INFO:
			// log.Print(200)        //INFO:
			// log.Print(tt.reqFile) //INFO:
			// log.Print(300)                              //INFO:
			// log.Print(testutil.LoadFile(t, tt.reqFile)) //INFO:

			sut := PostTasks{
				Repository: &in_memory.TaskRepository{
					Tasks: map[model.TaskID]*model.Task{},
				},
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			res := w.Result()
			// log.Print(300)     //INFO:
			// log.Print(20, res) //INFO:
			testutil.AssertResponse(t, res, tt.want.status, testutil.LoadFile(t, tt.want.resFile))
		})
	}
}
