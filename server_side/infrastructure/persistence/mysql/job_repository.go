package mysql

import (
	"github.com/tsubaoza0901/echo-rest-api/domain/model"
	"github.com/tsubaoza0901/echo-rest-api/domain/repository"
	"github.com/tsubaoza0901/echo-rest-api/infrastructure/rdb"
	"gorm.io/gorm"
)

type jobRepository struct {
	Conn *gorm.DB
}

func NewJobRepository(Conn *gorm.DB) repository.JobRepository {
	return &jobRepository{Conn}
}

func (r *jobRepository) FetchByID(id uint) (*model.Job, error) {
	return r.fetchByID(nil, id, false)
}

func (r *jobRepository) fetchByID(tx *gorm.DB, id uint, isExclusionControl bool) (*model.Job, error) {
	job := &rdb.Job{}
	job.ID = id

	if tx == nil {
		tx = r.Conn
	}

	if isExclusionControl {
		tx = tx.Unscoped()
	}

	err := tx.Find(job).Error
	if err != nil {
		return nil, err
	}

	return convertJob(job), nil
}
