package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestJobHandler_GetJob(t *testing.T) {
	type testData struct {
		name         string
		request      string
		responseCode int
	}
	tests := []testData{
		{
			"正常",
			"/api/v1/job/:id",
			http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testJobHandlerGetJob(t, test.request, test.responseCode)
		})
	}
}

func testJobHandlerGetJob(t *testing.T, request string, responseCode int) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, request, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, GetJob(c)) {
		t.Log("RESPONSECODE:", rec.Code)
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
