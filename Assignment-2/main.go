package main

import (
	"dalas98/golang-lesson/Assignment-2/app/controller"
	"dalas98/golang-lesson/Assignment-2/app/repository"
	"dalas98/golang-lesson/Assignment-2/app/routers"
	"dalas98/golang-lesson/Assignment-2/config"
	"math/rand"
	"time"
)

// @tittle To Do List API
// @description ini adalah sample api spec untuk To Do List API
// @version 1.0
// @termOfService https://swagger.io/terms/
// @contact.name Yusuf Farhan Hasbullah
// @contact.email yusuf.hasbullah@misteraladin.com
// @host localhost:9000
// @BasePath /todos
func main() {
	rand.Seed(time.Now().UnixNano())

	err := config.ConnectionMysql()
	if err != nil {
		panic(err)
	}

	db := config.GetMysql()

	todoRepo := repository.NewTodoRepository(db)

	todoHandler := controller.NewTodoHandler(todoRepo)

	router := routers.NewRoute(todoHandler)

	router.CreateRoute().Run(":9000")

}
