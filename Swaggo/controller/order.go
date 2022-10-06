package controller

import (
	"dalas98/golang-lesson/Swaggo/models"
	"dalas98/golang-lesson/Swaggo/params"
	"dalas98/golang-lesson/Swaggo/views"
	"encoding/json"
	"net/http"
	"time"
)

// GetOrder godoc
// @Summary Get All Orders
// @Description Get All orders in detail
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} views.GetOrderSuccess
// @Failure 404 {object} views.GetOrderFailedNotFound
// @Router /orders [get]
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(views.Response{
		StatusCode: 200,
		Message:    "GET ORDERS SUCCESS",
		Payload: []models.Order{
			{
				ID:           1,
				CustomerName: "MNC B",
				ProductsId:   2,
				UserId:       1,
				CreatedAt:    time.Now(),
			},
		},
	})
}

// CreateOrder godoc
// @Summary Create a new Order
// @Description Create a new Order with input payload
// @Tags orders
// @Accept json
// @Produce json
// @Param order body params.CreateOrder true "Create order payload"
// @Param id path int true "path"
// @Param Authorization header string false "token"
// @Param X-Token header string false "X-Token"
// @Success 200 {object} views.GetOrderSuccess
// @Failure 404 {object} views.GetOrderFailedNotFound
// @Router /orders/{id} [post]
func CreateOrders(w http.ResponseWriter, r *http.Request) {
	var req params.CreateOrder
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(views.Response{
		StatusCode: 200,
		Message:    "CREATE_ORDER_SUCCESS",
		Payload:    req,
	})
}
