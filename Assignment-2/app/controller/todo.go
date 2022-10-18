package controller

import (
	"dalas98/golang-lesson/Assignment-2/app"
	"dalas98/golang-lesson/Assignment-2/app/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Repo app.ITodoRepository
}

func NewTodoHandler(repo app.ITodoRepository) *TodoHandler {
	return &TodoHandler{
		Repo: repo,
	}
}

// GetTodos godoc
// @Summary Get all Todo
// @Description Get all todo
// @Tags Todo List
// @Accept json
// @Produce json
// @Success 200 {object} transformer.GetTodoSuccess
// @Failure 502 {object} transformer.GetTodoFailed
// @Router /todos [get]
func (t *TodoHandler) GetTaskHandler(ctx *gin.Context) {
	todos, err := t.Repo.GetTask()
	if err != nil {
		baseRespondFailed(err, ctx)
		return
	}

	baseRespondSuccess(todos, ctx)

}

// CreateTodo godoc
// @Summary Create Todo List
// @Description Create Todo List
// @Tags Todo List
// @Accept json
// @Produce json
// @Param todo body request.TodoRequest true "Create todo payload"
// @Success 200 {object} transformer.GetTodoSuccess
// @Failure 502 {object} transformer.GetTodoFailed
// @Router /todos [post]
func (t *TodoHandler) CreateTaskHandler(ctx *gin.Context) {
	var request model.Todo
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		baseRespondFailed(err, ctx)
		return
	}

	err = t.Repo.CreateTask(&request)
	if err != nil {
		baseRespondFailed(err, ctx)
	}

	baseRespondSuccess(nil, ctx)

}

// UpdateTodo godoc
// @Summary Update Todo List
// @Description Update Todo List
// @Tags Todo List
// @Accept json
// @Produce json
// @Param id path int false "id todo"
// @Param todo body request.TodoRequest true "Update todo payload"
// @Success 200 {object} transformer.GetTodoSuccess
// @Failure 502 {object} transformer.GetTodoFailed
// @Router /todos/{id} [put]
func (t *TodoHandler) UpdateTaskHandler(ctx *gin.Context) {
	var request model.Todo
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		baseRespondFailed(err, ctx)
		return
	}

	id, bool := ctx.Params.Get("id")
	if !bool {
		baseRespondFailed(errors.New("failed to get parameter id"), ctx)
		return
	}

	err = t.Repo.UpdateTask(id, &request)
	if err != nil {
		baseRespondFailed(err, ctx)
		return
	}

	baseRespondSuccess(nil, ctx)
}

// DeleteTodo godoc
// @Summary Delete Todo List
// @Description Delete Todo List
// @Tags Todo List
// @Accept json
// @Produce json
// @Param id path int false "id todo"
// @Success 200 {object} transformer.GetTodoSuccess
// @Failure 502 {object} transformer.GetTodoFailed
// @Router /todos/{id} [delete]
func (t *TodoHandler) DeleteTaskHandler(ctx *gin.Context) {
	id, bool := ctx.Params.Get("id")
	if !bool {
		baseRespondFailed(errors.New("failed to get parameter id"), ctx)
		return
	}

	err := t.Repo.DeleteTask(id)
	if err != nil {
		baseRespondFailed(err, ctx)
		return
	}
	baseRespondSuccess(nil, ctx)
}

func baseRespondSuccess(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "success",
	})
}

func baseRespondFailed(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusBadGateway, gin.H{
		"error":   err.Error(),
		"message": "failed",
	})
}
