package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// // JobHandler ...
// type JobHandler interface {
// 	CreateJob()
// 	GetJob()
// 	GetJobs(c echo.Context) error
// 	UpdateJob()
// 	DeleteJob()
// }

// CreateJob ...
func CreateJob(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "You can CreateJob!! "+"Name:"+name+",Email:"+email)
}

// GetJob ...
func GetJob(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "You can GetJob!! "+id)
}

// GetJobs ...
func GetJobs(c echo.Context) error {
	userID := c.QueryParam("user_id")
	return c.String(http.StatusOK, "You can GetJobs!!"+userID)
}

// UpdateJob ...
func UpdateJob(c echo.Context) error {
	return c.String(http.StatusOK, "You can UpdateJob!!")
}

// DeleteJob ...
func DeleteJob(c echo.Context) error {
	return c.String(http.StatusOK, "You can DeleteJob!!")
}
