package views

import (
	"dalas98/golang-lesson/Swaggo/models"
)

type GetOrderSuccess struct {
	StatusCode int            `json:"status" example:"200"`
	Message    string         `json:"message" example:"GET_SUCCESS"`
	Payload    []models.Order `json:"payload"`
}

type GetOrderFailedNotFound struct {
	StatusCode int    `json:"status" example:"404"`
	Message    string `json:"message" example:"GET_NOT_FOUND"`
}
