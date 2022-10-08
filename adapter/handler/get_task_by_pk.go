package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/repository"
	"github.com/teru-0529/go_webapi_sandbox/domain/service"
)

type GetTaskByPk struct {
	Repository *in_memory.TaskRepository
	Validator  *validator.Validate
}

func (gt *GetTaskByPk) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	p := struct {
		Id string `validate:"required,numeric,gt=0"`
	}{Id: chi.URLParam(r, "id")}

	// バリデーション（400エラー）
	err := gt.Validator.Struct(p)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(p.Id)
	service := service.GetTaskByPkService(gt.Repository, model.TaskID(id))

	// 実行
	err = service.Execute()

	if err != nil {
		if err == repository.ErrNotFound {
			RespondJSON(ctx, w, &ErrResponse{
				Message: err.Error(),
			}, http.StatusNotFound)
		} else {
			RespondJSON(ctx, w, &ErrResponse{
				Message: err.Error(),
			}, http.StatusInternalServerError)
		}
		return
	}

	res := service.Task
	RespondJSON(ctx, w, res, http.StatusOK)
}
