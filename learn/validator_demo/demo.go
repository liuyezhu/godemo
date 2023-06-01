package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

type Usersa struct {
	Name   string `form:"name" json:"name" validate:"required" label:"用dd户名"`
	Age    uint8  `form:"age" json:"age" validate:"required,gt=18" label:"年龄"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" validate:"required,len=6"`
}
type Users struct {
	Name   string `form:"name" json:"name" validate:"required" label:"用户名"`
	Age    uint8  `form:"age" json:"age" validate:"required,gt=18" label:"年龄"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" validate:"required,len=6"`
}

// 结构体类型
type Struct struct {
}

// 调用器接口
type Invoker interface {
	// 需要实现一个Call()方法
	Call(interface{})
}

// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

func main() {
	users := &Users{
		//Name:   "admin",
		Age:    12,
		Passwd: "123",
		Code:   "123456",
	}
	fmt.Println(users)
	usersa := Usersa{}
	fmt.Println(usersa)
	fmt.Println(reflect.TypeOf(usersa))
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		fmt.Println(name)
		return name
	})
	//注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usersa)
	err = validate.Struct(usersa)
	if err != nil {
		fmt.Println(err.(validator.ValidationErrors).Translate(trans))
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans))
			return
		}
	}

	return
}
