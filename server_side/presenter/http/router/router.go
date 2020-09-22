package router

import (
	"github.com/labstack/echo"
	"github.com/tsubaoza0901/echo-rest-api/presenter/http/handler"
)

// InitRouting ...
func InitRouting(e *echo.Echo) {
	apiBaseURL := "/api/v1/"

	// Job
	e.POST(apiBaseURL+"job", handler.CreateJob)
	e.GET(apiBaseURL+"job/:id", handler.GetJob)
	e.GET(apiBaseURL+"jobs", handler.GetJobs)
	e.PUT(apiBaseURL+"job/:id", handler.UpdateJob)
	e.DELETE(apiBaseURL+"job/:id", handler.DeleteJob)
}
