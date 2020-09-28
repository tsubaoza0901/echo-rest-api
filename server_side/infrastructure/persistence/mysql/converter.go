package mysql

import (
	"github.com/tsubaoza0901/echo-rest-api/domain/model"
	"github.com/tsubaoza0901/echo-rest-api/infrastructure/rdb"
)

func convertJob(jobRdb *rdb.Job) *model.Job {
	return &model.Job{
		ID:        jobRdb.ID,
		JobTitle:  jobRdb.JobTitle,
		UpdatedAt: jobRdb.UpdatedAt,
	}
}

func convertJobs(jobsRdb []*rdb.Job) []*model.Job {
	ret := make([]*model.Job, len(jobsRdb))
	for i, v := range jobsRdb {
		ret[i] = &model.Job{
			ID:        v.ID,
			JobTitle:  v.JobTitle,
			UpdatedAt: v.UpdatedAt,
		}
	}
	return ret
}
