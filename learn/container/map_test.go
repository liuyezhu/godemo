package main

import (
	"fmt"
	"testing"
)

//map创建
//make(map[string]int)
//map遍历
//使用range遍历key,遍历key,value对

//map的key
//除了slice,map,function的内建类型都可以作为key
//struct 类型不包含上述字段，也可以作为key
func TestMap(t *testing.T) {
	m := map[string]string{
		"name":    "lyz",
		"course":  "golang",
		"site":    "11",
		"quality": "o",
	}
	t.Log("定义map")
	m["aaa"] = "sddsd"
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m, m2, m3)
	//遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	if a, ok := m["lyz"]; ok {
		fmt.Println(a)
	} else {
		fmt.Println("key does not exist")
	}

	t.Log("删除map,key值")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	m4 := make(map[string]string)
	m4["test"] = "test"
	fmt.Println(m4)
}
