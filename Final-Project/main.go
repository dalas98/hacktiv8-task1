package main

import (
	"log"

	"dalas98/golang-lesson/Final-Project/server"

	_ "github.com/joho/godotenv/autoload"
)

// @title Hacktiv8 Mygram API
// @version 1.0
// @description MyGram API For Final Project Hacktiv8

// @BasePath /

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
