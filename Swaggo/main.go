package main

import (
	"dalas98/golang-lesson/Swaggo/controller"
	"net/http"

	_ "dalas98/golang-lesson/Swaggo/docs"

	httpSwagger "github.com/swaggo/swag"

	"github.com/gorilla/mux"
)

// @title Orders API
// @description Ini adalah sample API Spec untuk Api orders
// @version 1.0
// @termsOfService https://swagger.io/terms/
// @contact.Name Hacktiv8
// @contact.Email admin@hacktiv8.com
// @host localhost:4000
// @BasePath /v1
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/v1/orders", controller.GetOrders)

	router.PathPrefix("swagger", httpSwagger.WrapHandler)
	http.ListenAndServe(":4000", router)
}
