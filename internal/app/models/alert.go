package models

import "time"

type Alert struct {
	ID uint64
	time.Time
	Members []*User
}

type CreateAlert struct {
	time.Time
	Members []*User
}

type UpdateAlert struct {
	time.Time
	Members []*User
}
