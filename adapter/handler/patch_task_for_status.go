package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
	"github.com/teru-0529/go_webapi_sandbox/domain/service"
)

type PatchTaskForStatus struct {
	Repository *in_memory.TaskRepository
	Validator  *validator.Validate
}

func (pts *PatchTaskForStatus) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var p struct {
		Id string `validate:"required,numeric,gt=0"`
	}

	var b struct {
		Status string `json:"status" validate:"required,oneof=todo doing done"`
	}

	// デコード失敗（500エラー）
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// バリデーション（400エラー）
	err := pts.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	p.Id = chi.URLParam(r, "id")

	// バリデーション（400エラー）
	err = pts.Validator.Struct(p)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(p.Id)
	service := service.PatchTaskForStatusService(pts.Repository, model.TaskID(id), model.TaskStatus(b.Status))

	// 実行
	err = service.Execute()

	if err != nil {
		if err == in_memory.ErrNotFound {
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
