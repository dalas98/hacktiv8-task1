package models

type Payload struct {
	Status Status
}

type Status struct {
	Water int
	Wind  int
}
