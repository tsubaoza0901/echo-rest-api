package mock

import (
	"time"

	"github.com/tsubaoza0901/echo-rest-api/domain/model"
)

func getMockJob() *model.Job {
	i := &model.Job{
		ID:        1,
		JobTitle:  "案件タイトル",
		UpdatedAt: time.Date(2020, 3, 3, 17, 44, 13, 0, time.Local),
	}
	return i
}

// GetMockJob ...
func GetMockJob(option ...OptionJob) *model.Job {
	i := getMockJob()
	for _, setter := range option {
		setter(i)
	}
	return i
}

// OptionJob ...
type OptionJob func(b *model.Job)

// GetMockJobs ...
func GetMockJobs() []*model.Job {
	i := getMockJob()
	return []*model.Job{i, i}
}
