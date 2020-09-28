package usecase

import (
	"github.com/pkg/errors"
	"github.com/tsubaoza0901/echo-rest-api/domain/model"
	"github.com/tsubaoza0901/echo-rest-api/domain/repository"
)

// JobUseCase JobUseCase interface
// 役割：job_usecase.goに定義されたメソッドの呼び出しリスト
type JobUseCase interface {
	GetJob(id uint) (*model.Job, error)
	GetJobs(req *model.JobSearchRequest) ([]*model.Job, error)
}

// jobUseCase jobUseCase構造体
// 役割：埋め込んだinterfaceに定義されたメソッドを自身の構造体のメソッドとして取得
// 参考：https://qiita.com/tenntenn/items/bba69d84a1e0b67c4db8
// 参考：https://qiita.com/momotaro98/items/4f6e2facc40a3f37c3c3
type jobUseCase struct {
	repository.AppRepository
	repository.JobRepository
}

// NewJobUseCase NewJobUseCase関数
// 役割：jobUseCaseのコンストラクタ関数
func NewJobUseCase(app repository.AppRepository, jobr repository.JobRepository) JobUseCase {
	return &jobUseCase{app, jobr}
}

// GetJob GetJobメソッド
// 役割：指定したIDの案件情報取得
func (u *jobUseCase) GetJob(id uint) (*model.Job, error) {
	job, err := u.JobRepository.FetchByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "faild to FetchByID()")
	}
	return job, nil
}

// GetJobs GetJobsメソッド
// 役割：案件情報取得
func (u *jobUseCase) GetJobs(req *model.JobSearchRequest) ([]*model.Job, error) {
	jobs, err := u.JobRepository.Search(req)
	if err != nil {
		return nil, errors.Wrap(err, "faild to Search()")
	}
	return jobs, nil
}
