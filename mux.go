package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/teru-0529/go_webapi_sandbox/adapter/handler"
	"github.com/teru-0529/go_webapi_sandbox/adapter/repository/in_memory"
)

// コンストラクタ
func NewMux() http.Handler {
	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	v := validator.New()
	mux.Route("/tasks", func(r chi.Router) {
		pt := &handler.PostTasks{
			Repository: in_memory.InMemoryRepo,
			Validator:  v,
		}
		r.Post("/", pt.ServeHTTP)

		lt := &handler.ListTasks{
			Repository: in_memory.InMemoryRepo,
		}
		r.Get("/", lt.ServeHTTP)
	})
	return mux
}
