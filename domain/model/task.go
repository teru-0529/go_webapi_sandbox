package model

import "time"

type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID        TaskID     `json:"id"`
	Title     string     `json:"title"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Tasks []*Task

func NewTask(title string) *Task {
	task := &Task{
		Title:     title,
		Status:    TaskStatusTodo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return task
}
