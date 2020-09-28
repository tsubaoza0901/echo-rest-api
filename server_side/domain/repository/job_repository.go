package repository

import "github.com/tsubaoza0901/echo-rest-api/domain/model"

// JobRepository JobRepository構造体
// 役割：mysql/job_repository.goに定義されたメソッドの呼び出しリスト
type JobRepository interface {
	FetchByID(id uint) (*model.Job, error)
	Search(req *model.JobSearchRequest) ([]*model.Job, error)
}
