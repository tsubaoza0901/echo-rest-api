package rdb

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	JobTitle string `grom:"job_title; not null;"`
}
