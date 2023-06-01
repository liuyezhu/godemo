package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LoginParam struct {
	Account  string `validate:"required,len=6,contains=liu,custom_verify=VerifyUserUnique" form:"account" json:"account" label:"账号"`
	Password string `validate:"required" form:"password" json:"password" label:"密码"`
}

type RegParam struct {
	Account  string `validate:"required" form:"account" json:"account" label:"账号"`
	Password string `validate:"required" form:"password"  json:"password" label:"密码"`
}
type OrderFilterType uint

const (
	OrderFilterType_OFTNone        OrderFilterType = 0
	OrderFilterType_OFTVaild       OrderFilterType = 1
	OrderFilterType_OFTVaildCancel OrderFilterType = 2
	OrderFilterType_OFTCategory    OrderFilterType = 4
)

func main() {

	fmt.Println(OrderFilterType(7))
	router := gin.Default()
	fmt.Println(1 & 5689)
	router.POST("login", GetHandlerFunc(LoginHandlerFunc))
	router.POST("reg", GetHandlerFunc(RegHandlerFunc))
	_ = router.Run(":8882")

}

func LoginHandlerFunc(ctx *gin.Context, params *LoginParam) {
	fmt.Println(111)
	fmt.Println("这是参数", params)
	//todo 登录逻辑操作
}

func RegHandlerFunc(ctx *gin.Context) {
	fmt.Println("sdsd")
}

func (that *LoginParam) DeserializeFrom(ctx *gin.Context) error {
	return ctx.ShouldBindWith(that, binding.FormPost)
}

func (that *RegParam) DeserializeFrom1(ctx *gin.Context) error {
	return ctx.ShouldBindWith(that, binding.FormPost)
}

func (that *LoginParam) AAAA(ctx *gin.Context) error {
	return ctx.ShouldBindWith(that, binding.FormPost)
}
