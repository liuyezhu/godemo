package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type LoginParam struct {
	Account  string `validate:"required" json:"account"`
	Password string `validate:"required" json:"password"`
}

func main() {
	router := gin.Default()
	router.POST("login", LoginHandlerFunc)
	router.POST("post", func(c *gin.Context) {
		// 返回的 code 和 字符串返回
		c.String(200, "这是一个 post 方法\n")
	})
	_ = router.Run(":8881")

}

// HanlderFunc
func LoginHandlerFunc(ctx *gin.Context) {
	param := LoginParam{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		ctx.String(400, "请求参数异常")
		return
	}
	secc, msg := Validate(&param)
	if !secc {
		ctx.String(400, msg)
		return
	}
	// ... CRUD
}

func Validate(param interface{}) (bool, string) {
	validate := validator.New()
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	err:= zh_translations.RegisterDefaultTranslations(validate, trans)

	err = validate.Struct(param)
	if err != nil {
		fmt.Println(err)
	}
	errs := err.(validator.ValidationErrors)

	if len(errs) > 0 {
		err := errs[0]
		return false, err.Translate(trans)
	}
	return true, ""
}
