package models

type User struct {
	ID     uint64
	Name   string
	Status string
}

type CreateUser struct {
	Name   string
	Status string
}

type UpdateUser struct {
	Name   string
	Status string
}
