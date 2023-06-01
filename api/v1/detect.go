package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test_go_mod/response"
	"test_go_mod/service"
)

func Detect(c *gin.Context) {
	//responseBody := response.NewResponseBody()
	//defer response.RecoverResponse(ctx,responseBody)
	//param := &service.DetectRequestParams{Ctx: ctx}
	//users := service.Detect(param)
	//responseBody.SetData(users)
	//responseBody.SetErrNo(1)
	//responseBody.SetErrMsg("ss")

	//param := &service.DetectRequestParams{Ctx: ctx}
	//UserLists := service.Detect(param)
	//response.Success(ctx, UserLists, "12345")

	fmt.Println(111)
	param := &service.DetectRequestParams{Ctx: c}
	UserLists := service.Detect(param)
	//response.Error(c, 1, "ss", UserLists)
	//response.Err.ToString()
	c.JSON(http.StatusOK, response.OK.WithData(UserLists))
	c.JSON(http.StatusInternalServerError, response.ERROR.WithMsg("通用错误"))
	return
}

func Resp(c *gin.Context) {
	//response.Suc(c, response.WithData(dao.GetUser()))
	response.Suc(c, response.WithMsg("调用成功"))
}
