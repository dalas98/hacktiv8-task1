package models

import "time"

type Order struct {
	ID           int       `example:"1"`
	CustomerName string    `example:"MNC GOLANG LESSON"`
	ProductsId   int       `example:"1"`
	UserId       int       `example:"1"`
	CreatedAt    time.Time `exmaple:"time"`
}
