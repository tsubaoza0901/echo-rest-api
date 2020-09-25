package repository

import "github.com/tsubaoza0901/echo-rest-api/domain/model"

type JobRepository interface {
	FetchByID(id uint) (*model.Job, error)
}
