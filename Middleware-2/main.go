package main

import (
	"dalas98/golang-lesson/Middleware-2/server"
	"dalas98/golang-lesson/Middleware-2/server/controller"
)

func main() {

	userController := controller.NewUserController()

	router := server.NewGinRotuer(":4444", userController)

	router.Start()
}
