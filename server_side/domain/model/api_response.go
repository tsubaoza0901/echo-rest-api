package model

type APIResponse struct {
	Status   string      `json:"status" example:"A400"`
	Message  string      `json:"msg" example:"some error"`
	Response interface{} `json:"response"`
}

const (
	StatusSuccess       = "A200"
	StatusBadRequestErr = "A400"
	StatusNotFoundErr   = "A404"
	StatusServerErr     = "A500"
	StatusUnauthorized  = "A401"
)

var statusText = map[string]string{
	StatusSuccess: "Success",
}

func StatusText(code string) string {
	return statusText[code]
}

func NewAPIResponse(errCode int, msg string, res interface{}) *APIResponse {
	sts := StatusSuccess
	switch errCode {
	case ErrBadRequest:
		sts = StatusBadRequestErr
	case ErrRecordNotFound:
		sts = StatusNotFoundErr
	case ErrFailedToServer:
		sts = StatusServerErr
	case ErrUnautherized:
		sts = StatusUnauthorized
	}

	return &APIResponse{Status: sts, Message: msg, Response: res}
}
