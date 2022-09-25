package hander

import (
	"net/http"

	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
	"github.com/teru-0529/go_webapi_sandbox/domain/model"
)

type ListTasks struct {
	Repository *in_memory.TaskRepository
}

type task struct {
	ID     model.TaskID     `json:"id"`
	Title  string           `json:"title"`
	Status model.TaskStatus `json:"status"`
}

func (lt *ListTasks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// var taskRepo repository.TaskRepositorier = in_memory.NewTaskRepo() //FIXME:
	// tasks := taskRepo.List() //FIXME:
	tasks := lt.Repository.List()

	res := []task{}
	for _, t := range tasks {
		res = append(res, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}

	RespondJSON(ctx, w, res, http.StatusOK)
}
