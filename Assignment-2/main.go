package main

import (
	"dalas98/golang-lesson/Assignment-2/config"
	"fmt"
)

// @tittle To Do List API
// @description ini adalah sample api spec untuk To Do List API
// @version 1.0
// @termOfService https://swagger.io/terms/
// @contact.name Yusuf Farhan Hasbullah
// @contact.email yusuf.hasbullah@misteraladin.com
// @host localhost:4000
// @BasePath /
func main() {

	err := config.ConnectionMysql()
	if err != nil {
		panic(err)
	}

	db := config.GetMysql()
	if db == nil {
		panic("Can't Connect DB")
	}

	fmt.Println("DB Connected")
	defer db.Close()

}
