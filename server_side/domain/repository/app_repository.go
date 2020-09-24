package repository

import (
	"gorm.io/gorm"
)

type AppRepository interface {
	Begin() *gorm.DB
	Rollback(tx *gorm.DB) *gorm.DB
	Commit(tx *gorm.DB) *gorm.DB
}
