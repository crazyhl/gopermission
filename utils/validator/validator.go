package validator

import (
	"github.com/crazyhl/gopermission/v1/base_struct"
	"github.com/crazyhl/gopermission/v1/utils/trans"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
)

// 校验用户
func Validate(model interface{}) []error {
	var validateErrors []error
	validate := validator.New()
	err := validate.Struct(model)
	ts := trans.ZhTranslator()
	_ = zh.RegisterDefaultTranslations(validate, ts)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validateErrors = append(validateErrors, base_struct.ValidateError{
				Message: err.Translate(ts),
			})
		}
	}
	return validateErrors
}
