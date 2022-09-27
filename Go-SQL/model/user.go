package model

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Contact   string
	CreatedAt time.Time
}
