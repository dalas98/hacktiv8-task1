package app

import "dalas98/golang-lesson/Assignment-2/app/model"

type ITodoRepository interface {
	CreateTask(task *model.Todo) error
	GetTask() (*[]model.Todo, error)
	UpdateTask(id string, task *model.Todo) error
	DeleteTask(id string) error
}
