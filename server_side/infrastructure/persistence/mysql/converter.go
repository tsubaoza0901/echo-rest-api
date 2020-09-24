package mysql

import (
	"github.com/tsubaoza0901/echo-rest-api/domain/model"
	"github.com/tsubaoza0901/echo-rest-api/infrastructure/rdb"
)

func convertJob(jobrdb *rdb.Job) *model.Job {
	return &model.Job{
		ID:        jobrdb.ID,
		JobTitle:  jobrdb.JobTitle,
		UpdatedAt: jobrdb.UpdatedAt,
	}
}
