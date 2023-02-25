package models

import "time"

type Task struct {
	ID          uint64
	Name        string
	Description string
	CreateBy    *User
	Begin       time.Time
	End         time.Time
}

type CreateTask struct {
	Name        string
	Description string
	CreateBy    *User
	Begin       time.Time
	End         time.Time
}

type UpdateTask struct {
	Name        string
	Description string
	Begin       time.Time
	End         time.Time
}
