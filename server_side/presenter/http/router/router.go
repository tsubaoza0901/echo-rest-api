package router

import (
	"github.com/labstack/echo"
	"github.com/tsubaoza0901/echo-rest-api/presenter/http/handler"
)

// InitRouting ...
func InitRouting(e *echo.Echo, h handler.AppHandler) {
	apiBaseURL := "/api/v1/"

	// Job
	e.POST(apiBaseURL+"job", h.CreateJob)
	e.GET(apiBaseURL+"job/:id", h.GetJob)
	e.GET(apiBaseURL+"jobs", h.GetJobs)
	// e.PUT(apiBaseURL+"job/:id", h.UpdateJob)
	// e.DELETE(apiBaseURL+"job/:id", h.DeleteJob)
}
