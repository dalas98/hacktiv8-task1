package routers

import (
	"dalas98/golang-lesson/Assignment-2/app/controller"

	"github.com/gin-gonic/gin"

	docs "dalas98/golang-lesson/Assignment-2/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Todo *controller.TodoHandler
}

func NewRoute(todo *controller.TodoHandler) *Router {
	return &Router{
		Todo: todo,
	}
}

func (r *Router) CreateRoute() *gin.Engine {
	router := gin.Default()
	router.GET("todos", r.Todo.GetTaskHandler)
	router.POST("todos", r.Todo.CreateTaskHandler)
	router.PUT("todos/:id", r.Todo.UpdateTaskHandler)
	router.DELETE("todos/:id", r.Todo.DeleteTaskHandler)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
