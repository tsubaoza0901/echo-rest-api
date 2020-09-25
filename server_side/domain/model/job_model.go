package model

import "time"

type Job struct {
	ID        uint
	JobTitle  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
