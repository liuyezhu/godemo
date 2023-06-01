package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test_go_mod/response"
	"test_go_mod/utils/helper"
)

type InputRegData struct {
	Username string `form:"username" json:"username" validate:"required"`
	PassWord string `form:"password" json:"password,omitempty" validate:"required"`
	Email    string `form:"email" json:"email,omitempty" validate:"required"`
}

func Login(c *gin.Context) {

}

func Register(c *gin.Context,data * InputRegData) {
	fmt.Println(223)
	fmt.Println(data)
	return
	//fmt.Println(reflect.TypeOf(0))
	var regData = InputRegData{}
	err := c.ShouldBind(&regData)
	fmt.Println(regData)
	if err != nil {
		fmt.Println(err)
		response.Err(c, response.WithMsg("请求参数异常"))
		return
	}
	errMsg := helper.ValidatorStructParams(regData)
	fmt.Println(errMsg)
	return
	//errMsgA := helper.ValidatorStruct(regData)
	//fmt.Println("11111111111111111111")
	//fmt.Println(errMsgA)
	//return
	//fmt.Println(22)
	if errMsg != "" {
		response.Err(c, response.WithMsg(errMsg))
	}
	//if err := c.ShouldBind(&regData); err != nil {
	//	fmt.Println(regData)
	//	fmt.Println(reflect.TypeOf(regData))
	//	errMsg := helper.ValidatorStructParams(regData)
	//	fmt.Println(errMsg)
	//	fmt.Println(22)
	//	if errMsg != "" {
	//		response.Err(c, response.WithMsg(errMsg))
	//	}
	//}
	//if reflect.DeepEqual(regData, InputRegData{}) {
	//	response.Err(c)
	//	return
	//}

	//helper.VerifyUsername("454555445是多少")
	//dao.CreateUser(models.User{Username: regData.Username, Password: regData.PassWord})

	//saveData := &models.User{Username: regData.Username, Password: regData.PassWord}
	//saveErr := global.DB.Create(saveData)
	//fmt.Println(saveErr)
	//if saveErr != nil {
	//	fmt.Println(saveErr.Error)
	//}
	//fmt.Println(saveErr)
	//fmt.Println(22)
	response.Suc(c)
	return
}
