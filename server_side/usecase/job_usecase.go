package usecase

import (
	"github.com/pkg/errors"
	"github.com/tsubaoza0901/echo-rest-api/domain/model"
	"github.com/tsubaoza0901/echo-rest-api/domain/repository"
)

type JobUseCase interface {
	GetJob(id uint) (*model.Job, error)
}

type jobUseCase struct {
	repository.AppRepository
	repository.JobRepository
}

func NewJobUseCase(app repository.AppRepository, jobr repository.JobRepository) JobUseCase {
	return &jobUseCase{app, jobr}
}

func (o *jobUseCase) GetJob(id uint) (*model.Job, error) {
	job, err := o.JobRepository.FetchByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "faild to FetchByID()")
	}
	return job, nil
}
