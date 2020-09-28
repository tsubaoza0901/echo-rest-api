package rdb

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	JobTitle string `grom:"job_title; not null;"`
}

// TODO:プロパティ？の書き方がgromに依存するため、たしかにここでstructを定義する必要があるが、modelの内容と重複するため他にやり方がないか要検討
