package interactor

import (
	"github.com/tsubaoza0901/echo-rest-api/domain/repository"
	"github.com/tsubaoza0901/echo-rest-api/infrastructure/persistence/mysql"
	"github.com/tsubaoza0901/echo-rest-api/presenter/http/handler"
	"github.com/tsubaoza0901/echo-rest-api/usecase"
	"gorm.io/gorm"
)

// Interactor 安易コンテナ
type Interactor interface {
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Conn *gorm.DB
}

// NewInteractor interactorを取得
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

func (i *interactor) NewJobHandler() handler.JobHandler {
	return handler.NewJobHandler(i.NewJobUseCase())
}

func (i *interactor) NewJobUseCase() usecase.JobUseCase {
	return usecase.NewJobUseCase(
		i.NewAppRepository(),
		i.NewJobRepository(),
	)
}

func (i *interactor) NewAppRepository() repository.AppRepository {
	return mysql.NewAppRepository(i.Conn)
}

func (i *interactor) NewJobRepository() repository.JobRepository {
	return mysql.NewJobRepository(i.Conn)
}
