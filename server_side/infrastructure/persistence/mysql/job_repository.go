package mysql

import (
	"github.com/pkg/errors"
	"github.com/tsubaoza0901/echo-rest-api/domain/model"
	"github.com/tsubaoza0901/echo-rest-api/domain/repository"
	"github.com/tsubaoza0901/echo-rest-api/infrastructure/rdb"
	"gorm.io/gorm"
)

// jobRepository jobRepository構造体
// 役割：外部ORMのDBモデル
type jobRepository struct {
	Conn *gorm.DB
}

// NewJobRepository NewJobRepository関数
// 役割：jobRepositoryのコンストラクタ関数
func NewJobRepository(Conn *gorm.DB) repository.JobRepository {
	return &jobRepository{Conn}
}

// FetchByID FetchByIDメソッド
// 役割：指定されたIDに対応する単一レコードの取得
func (r *jobRepository) FetchByID(id uint) (*model.Job, error) {
	return r.fetchByID(nil, id, false)
}

// fetchByID fetchByIDメソッド
// 役割：
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

func (r *jobRepository) Search(req *model.JobSearchRequest) ([]*model.Job, error) {
	jobs := []*rdb.Job{}

	db := r.Conn
	
	if len(req.JobTitle) > 0 {
		db = db.Where("job_title LIKE ?", "%"+req.JobTitle+"%")
	}

	err := db.Order("id desc").Find(&jobs).Error
	if err != nil {
		err = model.NewAppError(model.ErrFailedToServer, errors.Wrap(err, "failed to jobs"))
		return nil, errors.WithStack(err)
	}
	return convertJobs(jobs), nil
}
