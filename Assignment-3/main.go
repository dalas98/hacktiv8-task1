package main

import (
	"dalas98/golang-lesson/Assignment-3/controller"
	"log"
	"net/http"
)

var port = ":9000"

func main() {
	// go forever()

	// quitChannel := make(chan os.Signal, 1)
	// signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	// <-quitChannel
	//time for cleanup before exit
	// fmt.Println("Stop")

	// http.HandleFunc("/write", controller.WriteToJson)
	http.HandleFunc("/load", controller.LoadJson)

	log.Println("http server runnin at port", port)
	http.ListenAndServe(port, nil)
}

// func forever() {
// 	for {
// 		go controller.WriteToJson()
// 		time.Sleep(time.Second * 15)
// 	}
// }
