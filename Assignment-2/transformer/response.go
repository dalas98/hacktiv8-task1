package transformer

import "dalas98/golang-lesson/Assignment-2/app/model"

type GetTodoSuccess struct {
	Data    *[]model.Todo `json:"data"`
	Message string        `json:"message" example:"success"`
}

type GetTodoFailed struct {
	Error   string `json:"error"`
	Message string `json:"message" example:"failed"`
}
