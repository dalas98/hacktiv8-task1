package controller

import (
	"dalas98/golang-lesson/Swaggo/models"
	"dalas98/golang-lesson/Swaggo/views"
	"encoding/json"
	"net/http"
	"time"
)

// GetOrder godoc
// @Summary Get All Orders
// @Description Get All orders in detail
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} views.GetOrderSuccess
// @Success 404 {object} views.GetOrderFailedNotFound
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
