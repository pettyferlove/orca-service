package util

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	trans    ut.Translator
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func InitTranslator(local string) (err error) {
	validate = validator.New()
	zhT := zh.New() //chinese
	enT := en.New() //english
	uni = ut.New(enT, zhT, enT)
	var o bool
	trans, o = uni.GetTranslator(local)
	if !o {
		return fmt.Errorf("uni.GetTranslator(%s) failed", local)
	}
	switch local {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(validate, trans)
	case "zh":
		err = chTranslations.RegisterDefaultTranslations(validate, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(validate, trans)
	}
	return

}

func Validate(param interface{}) (bool, string) {
	err := validate.Struct(param)
	var errs validator.ValidationErrors
	errors.As(err, &errs)
	if len(errs) > 0 {
		err := errs[0]
		return false, err.Translate(trans)
	}
	return true, ""
}
