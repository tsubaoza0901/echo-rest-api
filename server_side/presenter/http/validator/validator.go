package validator

// // メモ：バリデーションが必要な場合に実装
// // 参考：https://echo.labstack.com/guide/request
// // 参考：https://qiita.com/nanamen/items/c824a2c8f2e1767f90f8

// import (
// 	validator "github.com/go-playground/validator"
// 	"github.com/labstack/echo"
// )

// // CustomValidator CustomValidator構造体
// type CustomValidator struct {
// 	validator *validator.Validate
// }

// // NewValidator NewValidator
// func NewValidator() echo.Validator {
// 	customValidator := &CustomValidator{validator: validator.New()}
// 	_ = customValidator.validator.RegisterValidation("iframetag", customValidator.iframeTag)
// 	return customValidator
// }

// // Validate Validateメソッド
// // 役割：
// func (cv *CustomValidator) Validate(i interface{}) error {
// 	return cv.validator.Struct(i)
// }

// func (cv *CustomValidator) iframeTag(fl validator.FieldLevel) bool {
// 	return true
// }
