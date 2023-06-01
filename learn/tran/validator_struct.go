package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

type ControllerMapsType map[string]reflect.Value

//声明控制器函数Map类型变量
var ControllerMaps ControllerMapsType

//type Deserializer interface {
//	DeserializeFrom(ctx *gin.Context) error
//	DeserializeFrom1(ctx *gin.Context) error
//}

func ValidateC(param interface{}, c *gin.Context) (bool, string) {
	errBind := c.ShouldBind(param)
	if errBind != nil {
		return false, "请求参数异常"
	}
	validate := validator.New()

	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		if name == "" {
			name = fld.Tag.Get("json")
		}
		return name
	})
	fmt.Println(reflect.TypeOf(param).Method(0))

	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	err := translations.RegisterDefaultTranslations(validate, trans)
	err = validate.Struct(param)
	if err == nil {
		return true, ""
	}
	fmt.Println(err)
	errs := err.(validator.ValidationErrors)
	//todo 只返回一个错误
	if len(errs) > 0 {
		err := errs[0]
		fmt.Println("Namespace: ", err.Namespace())
		fmt.Println("Fild: ", err.Field())
		fmt.Println("StructNamespace: ", err.StructNamespace())
		fmt.Println("StructField: ", err.StructField())
		fmt.Println("Tag: ", err.Tag())
		fmt.Println("ActualTag: ", err.ActualTag())
		fmt.Println("Kind: ", err.Kind())
		fmt.Println("Type: ", err.Type())
		fmt.Println("Value: ", err.Value())
		fmt.Println("Param: ", err.Param())
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
	fmt.Println(param)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
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
	//denser, ok := param.(Deserializer)
	//// 如果未实现 Deserializer 接口则说明该 struct 使用默认绑定过程即可.
	//if !ok {
	//	fmt.Println(444)
	//	// 绑定参数到 struct
	//	err := ctx.ShouldBind(&param)
	//	fmt.Println(err)
	//	if err != nil {
	//		ctx.String(400, "请求参数异常")
	//		return
	//		//return false ,"请求参数异常"
	//	}
	//} else {
	//	// 绑定请求参数
	//	err := denser.DeserializeFrom(ctx)
	//	if err != nil {
	//		ctx.String(400, "请求参数异常")
	//		return
	//	}
	//	param = reflect.ValueOf(denser).Interface()
	//}
	// 验证参数
	suc, msg := ValidateC(param, ctx)
	if suc == false {
		//todo 校验失败逻辑
		ctx.String(400, msg)
		return
	}
	//todo 校验无错误 回调目标函数
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
