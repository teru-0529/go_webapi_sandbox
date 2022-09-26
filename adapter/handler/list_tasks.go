package handler

import (
	"net/http"

	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/service"
)

type ListTasks struct {
	Repository *in_memory.TaskRepository
}

func (lt *ListTasks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	service := service.ListTasksService(lt.Repository)

	// 実行
	_ = service.Execute()

	res := service.Tasks
	RespondJSON(ctx, w, res, http.StatusOK)
}
