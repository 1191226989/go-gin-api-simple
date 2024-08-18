package validation

import (
	"fmt"
	"reflect"

	"go-gin-api-simple/configs"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	lang := configs.Get().Language.Local
	validate := binding.Validator.Engine().(*validator.Validate)

	if lang == configs.ZhCN {
		trans, _ = ut.New(zh.New()).GetTranslator("zh")
		if err := zhTranslation.RegisterDefaultTranslations(validate, trans); err != nil {
			fmt.Println("validator zh translation error", err)
		}
		// 将错误提示的验证字段名改为中文名
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("label")
		})
	}

	if lang == configs.EnUS {
		trans, _ = ut.New(en.New()).GetTranslator("en")
		if err := enTranslation.RegisterDefaultTranslations(validate, trans); err != nil {
			fmt.Println("validator en translation error", err)
		}
	}
}

func Error(err error) (message string) {
	if validationErrors, ok := err.(validator.ValidationErrors); !ok {
		return err.Error()
	} else {
		for _, e := range validationErrors {
			message += e.Translate(trans) + ";"
		}
	}
	return message
}

// 自定义验证错误提示
func CustomErrorMessage(err error, req interface{}) string {
	request := reflect.TypeOf(req)
	if errs, ok := err.(validator.ValidationErrors); ok {

		for _, e := range errs {
			if f, exists := request.Elem().FieldByName(e.StructField()); exists {
				return f.Tag.Get("msg")
			}
		}
	}
	return err.Error()
}
