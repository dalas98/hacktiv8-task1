package models

import "time"

type Order struct {
	ID           int
	CustomerName string
	ProductsId   int
	UserId       int
	CreatedAt    time.Time
}
