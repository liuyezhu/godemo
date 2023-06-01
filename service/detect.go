package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DetectRequestParams struct {
	Text string `json:"text"`
	Ctx  *gin.Context
}

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Data struct {
	UserInfo []UserInfo `json:"Userinfo"`
}

func Detect(param *DetectRequestParams) []UserInfo {
	fmt.Println(param.Ctx.PostForm("id"))

	user1 := UserInfo{Name: "lyz-rrzuji", Age: 26}
	user2 := UserInfo{Name: "lyz-rrzuji", Age: 25}
	user3 := UserInfo{Name: "lyz-rrzuji", Age: 25}
	b, _ := json.Marshal(user1)
	//
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		fmt.Printf("key:%v value:%v\n", k, v)
	}
	list := Data{
		[]UserInfo{
			user1, user2, user3,
		},
	}
	for i := 1; i < 2; i++ {
		user := UserInfo{Name: "lyz-rrzuji" + strconv.Itoa(i), Age: 25}
		list.UserInfo = append(list.UserInfo, user)
	}
	return list.UserInfo
}
