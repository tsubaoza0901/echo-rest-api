package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/tsubaoza0901/echo-rest-api/conf"
	"github.com/tsubaoza0901/echo-rest-api/domain/model"
	"github.com/tsubaoza0901/echo-rest-api/usecase"
	"go.uber.org/zap"
)

// JobSearchRequest JobSearchRequest構造体
// 役割：
type JobSearchRequest struct {
	JobTitle string `query:"job_title"`
}

// JobResponse JobResponse構造体
// 役割：
type JobResponse struct {
	ID        uint      `json:"id" example:"1"`
	JobTitle  string    `json:"job_title" example:"案件タイトル"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-02-02T10:03:48.1292"`
}

// JobList JobList構造体
type JobList struct {
	Jobs []*JobResponse `json:"jobs"`
}

// JobHandler JobHandler interface
// 役割：job_handler.goに定義されたメソッドの呼び出しリスト
type JobHandler interface {
	GetJob(c echo.Context) error
	GetJobs(c echo.Context) error
	// CreateJob()
	// UpdateJob()
	// DeleteJob()
}

// jobHandler jobHandler構造体
// 役割：埋め込んだinterfaceに定義されたメソッドを自身の構造体のメソッドとして取得
type jobHandler struct {
	usecase.JobUseCase
}

// NewJobHandler NewJobHandler関数
// 役割：jobHandlerのコンストラクタ関数
func NewJobHandler(u usecase.JobUseCase) JobHandler {
	return &jobHandler{u}
}

// NewJobResponse NewJobResponse関数
// 役割：
func NewJobResponse(input *model.Job) *JobResponse {
	return &JobResponse{
		ID:        input.ID,
		JobTitle:  input.JobTitle,
		UpdatedAt: input.UpdatedAt,
	}
}

// NewJobList NewJobList関数
// 役割：
func NewJobList(input []*model.Job) *JobList {
	jobl := &JobList{
		Jobs: []*JobResponse{},
	}
	for _, job := range input {
		jobl.Jobs = append(jobl.Jobs, NewJobResponse(job))
	}
	return jobl
}

// GetJob GetJobメソッド
// 役割：
// @Resource /v1/job
// @Router api/v1/job/{id} [get]
func (h *jobHandler) GetJob(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	if err != nil || id <= 0 {
		err = model.NewAppError(model.ErrBadRequest, err).Wrap()
		return c.JSON(http.StatusOK, model.NewAPIResponse(model.ErrBadRequest, err.Error(), nil))
	}
	job, err := h.JobUseCase.GetJob(uint(id))
	if err != nil {
		code := model.ErrFailedToServer
		if ae, ok := errors.Cause(err).(*model.AppError); ok {
			code = ae.Code
			err = ae.Wrap()
		} else {
			err = model.NewAppError(model.ErrFailedToServer, err).Wrap()
		}
		return c.JSON(http.StatusOK, model.NewAPIResponse(code, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.NewAPIResponse(0, model.StatusText(model.StatusSuccess), NewJobResponse(job)))
}

// // GetJobs GetJobsメソッド
// 役割：
// @Resource /v1/jobs
// @Router api/v1/jobs [get]
func (h *jobHandler) GetJobs(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	req := &JobSearchRequest{}

	// bind req
	if err := c.Bind(req); err != nil {
		conf.WithContext(ctx).Debug("bind error", zap.Error(err))
		err = model.NewAppError(model.ErrBadRequest, err).Wrap()
		return c.JSON(http.StatusOK, model.NewAPIResponse(model.ErrBadRequest, err.Error(), nil))
	}

	// // メモ:今のままだと、ぬるぽエラーになるためコメントアウト。おそらく、validator.goの実装内容のため、バリデーションが必要な際に再編集。
	// // validate req
	// if err := c.Validate(req); err != nil {
	// 	conf.WithContext(ctx).Debug("validate error", zap.Error(err))
	// 	err = model.NewAppError(model.ErrBadRequest, err).Wrap()
	// 	return c.JSON(http.StatusOK, model.NewAPIResponse(model.ErrBadRequest, err.Error(), nil))
	// }

	request := &model.JobSearchRequest{
		JobTitle: req.JobTitle,
	}

	job, err := h.JobUseCase.GetJobs(request)
	if err != nil {
		code := model.ErrFailedToServer
		if ae, ok := errors.Cause(err).(*model.AppError); ok {
			code = ae.Code
			err = ae.Wrap()
		} else {
			err = model.NewAppError(model.ErrFailedToServer, err).Wrap()
		}
		return c.JSON(http.StatusOK, model.NewAPIResponse(code, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.NewAPIResponse(0, model.StatusText(model.StatusSuccess), NewJobList(job).Jobs))
}

// // CreateJob ...
// func CreateJob(c echo.Context) error {
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")
// 	return c.String(http.StatusOK, "You can CreateJob!! "+"Name:"+name+",Email:"+email)
// }

// // UpdateJob ...
// func UpdateJob(c echo.Context) error {
// 	return c.String(http.StatusOK, "You can UpdateJob!!")
// }

// // DeleteJob ...
// func DeleteJob(c echo.Context) error {
// 	return c.String(http.StatusOK, "You can DeleteJob!!")
// }
