package model

const (
	_ = iota
	ErrBadRequest
	ErrRecordNotFound
	ErrFailedToServer
	ErrUnautherized
)

var errorText = map[int]string{
	ErrBadRequest:     "不正な値が設定されています。",
	ErrRecordNotFound: "データが削除されているか存在しません。",
	ErrFailedToServer: "予期せぬエラーが発生しました。",
	ErrUnautherized:   "認証エラーが発生しました。",
}

func ErrorText(code int) string {
	return errorText[code]
}
