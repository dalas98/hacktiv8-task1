package main

import (
	"dalas98/golang-lesson/Go-SQL/config"
	"dalas98/golang-lesson/Go-SQL/faker"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	err := config.ConnectionPostgres()
	if err != nil {
		panic(err)
	}

	db := config.GetPostgres()
	if db == nil {
		panic("db not connect")
	}

	fmt.Println("db connect")
	defer db.Close()

	faker.CreateUser(2)
}
