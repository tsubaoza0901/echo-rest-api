package model

// APIResponse APIResponse構造体
// 役割：APIの呼び出し側に返却するResponse情報のモデル
type APIResponse struct {
	Status   string      `json:"status" example:"A400"`
	Message  string      `json:"msg" example:"some error"`
	Response interface{} `json:"response"`
}

// APIResponse Status
// 役割：APIResponseのStatusコード定義
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

// StatusText StatusText関数
// 役割：APIResponse Statusを元にしたAPI呼び出し側へのレスポンスメッセージ取得
func StatusText(status string) string {
	return statusText[status]
}

// NewAPIResponse NewAPIResponse関数
// 役割：処理結果を元にしたAPI呼び出し側へのレスポンス情報生成
func NewAPIResponse(errCode int, msg string, res interface{}) *APIResponse {
	sts := StatusSuccess
	switch errCode {
	case ErrBadRequest,
		ErrExistSameName,
		ErrUsedDesignTemplate,
		ErrExclusionControl,
		ErrUsedPlacementRelation,
		ErrStandardTemplate,
		ErrExistSameAdID,
		ErrUsedDeviceModel,
		ErrUsedDeviceModelGroup,
		ErrUsedOS,
		ErrUsedOSGroup,
		ErrUsedSDKVersion,
		ErrExistSameOuterDealID,
		ErrChangeOuterDealID,
		ErrUnsupportedPMP,
		ErrUsedDeal,
		ErrUnsupportedAutomaticRatio,
		ErrChangeAdFormat,
		ErrChangeDesignTemplateType,
		ErrChangeMediaType,
		ErrChangePlatform,
		ErrExistSameTemplateTypePlatformAdFormat,
		ErrNoPermission,
		ErrExistEncryptKeyPlatform:
		sts = StatusBadRequestErr
	case ErrRecordNotFound:
		sts = StatusNotFoundErr
	case ErrFailedToServer:
		sts = StatusServerErr
	case ErrUnauthorized:
		sts = StatusUnauthorized
	}

	return &APIResponse{Status: sts, Message: msg, Response: res}
}
