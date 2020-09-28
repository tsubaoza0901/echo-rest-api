package interactor

import (
	"github.com/tsubaoza0901/echo-rest-api/domain/repository"
	"github.com/tsubaoza0901/echo-rest-api/infrastructure/persistence/mysql"
	"github.com/tsubaoza0901/echo-rest-api/presenter/http/handler"
	"github.com/tsubaoza0901/echo-rest-api/usecase"
	"gorm.io/gorm"
)

// Interactor Interactor
// 役割：すべての依存関係を把握する設定ファイル（安易コンテナ）
type Interactor interface {
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Conn *gorm.DB
}

// NewInteractor NewInteractor関数
// 役割：interactorのコンストラクタ関数
func NewInteractor(Conn *gorm.DB) Interactor {
	return &interactor{Conn}
}

type appHandler struct {
	handler.JobHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{
		i.NewJobHandler(),
	}
	return appHandler
}

// NewJobHandler NewJobHandlerメソッド
// 役割：インスタンス生成
func (i *interactor) NewJobHandler() handler.JobHandler {
	// JobHandlerは、JobUseCaseに依存
	return handler.NewJobHandler(i.NewJobUseCase())
}

// NewJobUseCase NewJobUseCaseメソッド
// 役割：インスタンス生成
func (i *interactor) NewJobUseCase() usecase.JobUseCase {
	// JobUseCaseは、AppRepositoryとJobRepositoryに依存
	return usecase.NewJobUseCase(
		i.NewAppRepository(),
		i.NewJobRepository(),
	)
}

// NewAppRepository NewAppRepositoryメソッド
// 役割：インスタンス生成
func (i *interactor) NewAppRepository() repository.AppRepository {
	// AppRepositoryは、gormに依存
	return mysql.NewAppRepository(i.Conn)
}

// NewJobRepository NewJobRepositoryメソッド
// 役割：インスタンス生成
func (i *interactor) NewJobRepository() repository.JobRepository {
	// JobRepositoryは、gormに依存
	return mysql.NewJobRepository(i.Conn)
}
