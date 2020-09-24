package mysql

import (
	"github.com/tsubaoza0901/echo-rest-api/domain/repository"
	"gorm.io/gorm"
)

type appRepository struct {
	Conn *gorm.DB
}

func NewAppRepository(Conn *gorm.DB) repository.AppRepository {
	return &appRepository{Conn}
}

func (r *appRepository) Begin() *gorm.DB {
	return r.Conn.Begin()
}

func (r *appRepository) Rollback(tx *gorm.DB) *gorm.DB {
	return r.Conn.Rollback()
}

func (r *appRepository) Commit(tx *gorm.DB) *gorm.DB {
	return r.Conn.Commit()
}
