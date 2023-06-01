package main

import (
	"encoding/json"
	"fmt"
	"test_go_mod/global"
	"test_go_mod/models"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//type users struct {
//	Key int `json:"key"`
//}

func main() {
	err := global.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	user1 := user{}
	b, _ := json.Marshal(user1)
	u := make(map[string]interface{})
	_ = json.Unmarshal(b, &u)
	u["id"] = "11"
	u["pp"] = "11"
	u["sdad"] = "11dsdssd"
	fmt.Println(u["sdad"])

	arr := make(map[int]map[string]interface{})
	if arr[0] == nil {
		arr[0] = make(map[string]interface{})
		arr[0]["id"] = 1
		arr[0]["name"] = "lyz"
	}
	arr[23]["aa"] = "ass"
	fmt.Println(arr[23]["aa"])
	fmt.Println(arr[24])
	fmt.Println(arr[25])
	fmt.Println(arr[23])
	fmt.Println(arr[23])

	//m := make(map[string]map[string]string)
	//if m["user"] == nil {
	//	m["user"] = make(map[string]string)
	//	m["user"]["name"] = "lyz"
	//}
	//var a [1][4]int
	////a := [...]int{}
	//for i, v := range a {
	//	fmt.Println(i, v)
	//	fmt.Println(i, v)
	//}
	////fmt.Println(arr)
	//fmt.Println(m)
	//fmt.Println(a[0][3])
	var usersData []user
	fmt.Println(usersData)
	l, _ := json.Marshal(usersData)
	k := make(map[int]map[string]interface{})
	_ = json.Unmarshal(l, &k)
	//k[0]["ss"] = "ss"
	//fmt.Println(l)
	fmt.Println(k)
}
