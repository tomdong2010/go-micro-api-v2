/**
 * Created by Goland.
 * User: yan.wang5
 * Date: 2019/9/2
 */
package validator

import (
	"demo/cmn"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"

	"reflect"
	"regexp"
)

var Validate *validator.Validate
var translator ut.Translator

func init() {
	InitValidator()
}

// 初始化验证器
func InitValidator() {
	Validate = validator.New()
	Validate.RegisterTagNameFunc(validateShowCustomTag)

	_en := en.New()
	translator, _ = ut.New(_en, _en).GetTranslator("en")
	en_translations.RegisterDefaultTranslations(Validate, translator)

	registerValidator("username", "{0} is a invalid username", validateUserName)
	registerValidator("cc", "{0} is a invalid cc", validateCc)
	registerValidator("phone", "{0} is a invalid phone", validatePhone)
}

// 解析错误
func ParseErr(err error) string {
	if e, ok := err.(validator.ValidationErrors); ok {
		for _, v := range e {
			return v.Translate(translator)
		}
	}
	return err.Error()
}

func registerValidator(tag, msg string, f validator.Func) {
	Validate.RegisterValidation(tag, f)
	Validate.RegisterTranslation(tag, translator, func(ut ut.Translator) error {
		return ut.Add(tag, msg, false)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T(fe.Tag(), fe.Field())
		if err != nil {
			return fe.(error).Error()
		}
		return t
	})
}

// 验证后展示自定义的json字段
func validateShowCustomTag(fld reflect.StructField) string {
	return fld.Tag.Get("json")
}

// 用户名验证
func validateUserName(f validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(cmn.REGEXP_USERNAME, f.Field().String())
	return ok
}

// 国家码验证
func validateCc(f validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(cmn.REGEXP_CC, f.Field().String())
	return ok
}

// 手机号码验证
func validatePhone(f validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(cmn.REGEXP_PHONE, f.Field().String())
	return ok
}
