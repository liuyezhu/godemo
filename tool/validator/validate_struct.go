package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"test_go_mod/response"
)

func actionValidate(param interface{}, c *gin.Context) (bool, string) {
	errBind := c.ShouldBind(param)
	if errBind != nil {
		return false, "请求参数异常"
	}
	validate := validator.New()
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	err := translations.RegisterDefaultTranslations(validate, trans)
	err = validate.Struct(param)
	if err == nil {
		return true, ""
	}
	errs := err.(validator.ValidationErrors)
	if len(errs) > 0 {
		err := errs[0]
		errKey := fmt.Sprintf("%s.%s", err.StructNamespace(), err.ActualTag())
		if errMsg, ok := ErrMsgMap[errKey]; ok {
			return false, errMsg
		}
		return false, err.Translate(trans)
	}
	return true, ""
}

func valOf(i ...interface{}) []reflect.Value {
	var rt []reflect.Value
	for _, i2 := range i {
		rt = append(rt, reflect.ValueOf(i2))
	}
	return rt
}

func Validate(param interface{}) (bool, string) {
	validate := validator.New()
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	err := translations.RegisterDefaultTranslations(validate, trans)

	err = validate.Struct(param)
	if err == nil {
		return true, ""
	}
	errs := err.(validator.ValidationErrors)

	if len(errs) > 0 {
		err := errs[0]
		return false, err.Translate(trans)
	}
	return true, ""
}

func proxyHandlerFunc(ctx *gin.Context, handlerFunc interface{}) {
	funcType := reflect.TypeOf(handlerFunc)
	// 判断是否 Func
	if funcType.Kind() != reflect.Func {
		panic("the route handlerFunc must be a function")
	}
	typeParam := funcType.In(1).Elem()
	param := reflect.New(typeParam).Interface()
	// 验证参数
	suc, msg := actionValidate(param, ctx)
	if suc == false {
		//todo 校验失败逻辑
		response.Err(ctx, response.WithMsg(msg))
	}
	//成功回调目标方法
	valueFunc := reflect.ValueOf(handlerFunc)
	valueFunc.Call(valOf(ctx, param))
}

func GetHandlerFunc(handlerFunc interface{}) func(*gin.Context) {
	// 获取参数数量
	paramNum := reflect.TypeOf(handlerFunc).NumIn()
	valueFunc := reflect.ValueOf(handlerFunc)
	return func(context *gin.Context) {
		// 只有一个参数说明是未重构的 HandlerFunc
		if paramNum == 1 {
			valueFunc.Call(valOf(context))
			return
		}
		proxyHandlerFunc(context, handlerFunc)
	}
}
