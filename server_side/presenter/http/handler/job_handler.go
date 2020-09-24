package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/tsubaoza0901/echo-rest-api/domain/model"
	"github.com/tsubaoza0901/echo-rest-api/usecase"
)

type JobSearchRequest struct {
	JobTitle string `query:"job_title"`
}

type JobResponse struct {
	ID        uint      `json:"id" example:"1"`
	JobTitle  string    `json:"job_title" example:"案件タイトル"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-02-02T10:03:48.1292"`
}

// JobHandler ...
type JobHandler interface {
	// CreateJob()
	GetJob(c echo.Context) error
	// GetJobs(c echo.Context) error
	// UpdateJob()
	// DeleteJob()
}

type jobHandler struct {
	JobUseCase usecase.JobUseCase
}

func NewJobHandler(o usecase.JobUseCase) JobHandler {
	return &jobHandler{o}
}

func NewJobResponse(input *model.Job) *JobResponse {
	return &JobResponse{
		ID:        input.ID,
		JobTitle:  input.JobTitle,
		UpdatedAt: input.UpdatedAt,
	}
}

// // CreateJob ...
// func CreateJob(c echo.Context) error {
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")
// 	return c.String(http.StatusOK, "You can CreateJob!! "+"Name:"+name+",Email:"+email)
// }

// GetJob ...
func (h *jobHandler) GetJob(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return c.JSON(http.StatusOK, "NG")
	}
	job, err := h.JobUseCase.GetJob(uint(id))
	if err != nil {
		return c.JSON(http.StatusOK, "NG")
	}
	return c.JSON(http.StatusOK, model.NewAPIResponse(0, model.StatusText(model.StatusSuccess), NewJobResponse(job)))
}

// // GetJobs ...
// func GetJobs(c echo.Context) error {
// 	userID := c.QueryParam("user_id")
// 	return c.String(http.StatusOK, "You can GetJobs!!"+userID)
// }

// // UpdateJob ...
// func UpdateJob(c echo.Context) error {
// 	return c.String(http.StatusOK, "You can UpdateJob!!")
// }

// // DeleteJob ...
// func DeleteJob(c echo.Context) error {
// 	return c.String(http.StatusOK, "You can DeleteJob!!")
// }
