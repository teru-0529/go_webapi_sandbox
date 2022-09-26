package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/service"
)

type PostTasks struct {
	Repository *in_memory.TaskRepository
	Validator  *validator.Validate
}

func (pt *PostTasks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Title string `json:"title" validate:"required"`
	}

	// デコード失敗（500エラー）
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// バリデーション（400エラー）
	err := pt.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	task := model.NewTask(b.Title)
	service := service.PostTasksService(pt.Repository, task)

	// 実行
	err = service.Execute()

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}

	res := struct {
		ID model.TaskID `json:"id"`
	}{ID: service.Task.ID}
	RespondJSON(ctx, w, res, http.StatusCreated)
}
