package model

import "time"

type TaskID int64
type TaskStatus string

type UpdateField string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

const (
	F_Title  UpdateField = "Title"
	F_Status UpdateField = "Status"
)

type Task struct {
	ID         TaskID     `json:"id"`
	Title      string     `json:"title"`
	Status     TaskStatus `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt time.Time  `json:"modified_at"`
}

type Tasks []*Task

func NewTask(title string) *Task {
	task := &Task{
		Title:  title,
		Status: TaskStatusTodo,
	}
	return task
}
