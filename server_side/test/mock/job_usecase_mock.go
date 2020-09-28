package mock

import (
	"github.com/pkg/errors"
	"github.com/tsubaoza0901/echo-rest-api/domain/model"
)

// JobUseCase Mock JobUseCase
type JobUseCase struct{}

var (
	// IsGetJobErr GetJobエラーフラグ
	IsGetJobErr  bool
	IsGetJobsErr bool
)

func (j *JobUseCase) GetJob(id uint) (*model.Job, error) {
	if IsGetJobErr {
		return nil, errors.New("mock GetJob() err")
	}
	return GetMockJob(), nil
}

func (j *JobUseCase) GetJobs(req *model.JobSearchRequest) ([]*model.Job, error) {
	if IsGetJobsErr {
		return nil, errors.New("mock GetJobs() err")
	}
	return GetMockJobs(), nil
}
