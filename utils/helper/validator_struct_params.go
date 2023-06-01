package helper

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func ValidatorStructParams(params interface{}) string {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		if name == "" {
			name = fld.Tag.Get("json")
		}
		//fmt.Println(name)
		return name
	})
	//注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(params)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err.(validator.ValidationErrors).Translate(trans))
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans))
			return err.Translate(trans)
		}
	}
	return ""
}
