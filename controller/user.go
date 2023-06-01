package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"test_go_mod/dao"
	"test_go_mod/global"
	"test_go_mod/models"
	"test_go_mod/response"
)

type requestParams struct {
	page     int
	pageSize int
}

func GetUserList(c *gin.Context) {
	log.SetFlags(log.Ldate | log.Lshortfile)
	params := new(requestParams)
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		err = fmt.Errorf("got an error when getting the employee information: %v", err)
		//log.Fatalf(err.Error())
		log.Println(err.Error())

		//log.Panicln("日志信息3")
		//fmt.Println(err.Error())
		//fmt.Println(11)
	}
	params.pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "1"))
	fmt.Println(params)

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	var total, userList = dao.GetUserListDao(page, params.pageSize)
	// 判断
	if (total + len(userList)) == 0 {
		response.Error(c, 1, "", map[string]interface{}{
			"total":    total,
			"userList": userList,
		})
		return
	}
	response.Success(c, map[string]interface{}{
		"total":    total,
		"userList": userList,
	}, "")
}

func GetUser(ctx *gin.Context) {
	//defer func() {
	//	//var a interface{}
	//	if err := recover() ; err !=nil {
	//		fmt.Println(reflect.TypeOf(err)) // x的类型是int
	//		//response.Err(ctx,response.WithMsg())
	//	}
	//	fmt.Println(222)
	//}()
	user := dao.GetUser()
	fmt.Println(user)
	response.Suc(ctx, response.WithData(user))
}

//func CreateUser(c *gin.Context) {
//	dao.CreateUser()
//}
func Test(c *gin.Context) {
	err := global.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("失败")
	} else {
		fmt.Println("成功")
	}
}
