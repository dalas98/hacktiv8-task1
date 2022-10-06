package main

import (
	"dalas98/golang-lesson/Session-9/controller"
	"dalas98/golang-lesson/Session-9/httpclient"
	"log"
	"net/http"
)

var url = "https://jsonplaceholder.typicode.com/posts"
var port = ":4000"

func main() {
	client := httpclient.NewClient(10, url)

	postController := controller.NewPostController(client)

	http.HandleFunc("/posts", postController.GetPosts)
	http.HandleFunc("/posts/create", controller.Trace(controller.AllowOnlyPost(controller.Auth(postController.CreatePost))))

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}
