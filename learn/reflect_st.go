package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	ordId        int    `json:"order_id" validate:"required"`
	customerName string `json:"customerName" validate:"required"`
	callback     func() `json:"call_back" validate:"required"`
}

func InterfaceToStruct(object interface{}) {
	fmt.Println(reflect.TypeOf(object))
	obj := object.(Order)
	fmt.Println(88)
	fmt.Println(reflect.TypeOf(obj))
	fmt.Printf("Object = %v\n", obj)
}

// 反射传入的结构体并实例化一个实例，根据结构体字段在CSV行中找对应数据，找到的就设值
func instantiate(unknownSt interface{}, jsonData map[string]interface{}) interface{} {
	// 获取结构体的类型反射
	sType := reflect.TypeOf(unknownSt).Elem()
	// 根据类型实例化一个新结构体
	retSt := reflect.New(sType).Elem()
	// 遍历每一个结构体类型成员
	for i := 0; i < sType.NumField(); i++ {
		// 成员类型
		f := sType.Field(i)
		// 新结构体的对应成员
		fv := retSt.Field(i)
		// 查找数据Map，赋值
		if v, ok := jsonData[f.Tag.Get("json")]; ok && v != nil {
			fv.Set(reflect.ValueOf(v))
		}
	}
	return retSt
}

type tSt struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func main() {
	a := "aa"
	b := 6
	jd := make(map[string]interface{}, 0)
	jd["a"] = a
	jd["b"] = b
	ret := instantiate(&tSt{}, jd)
	fmt.Println(reflect.TypeOf(ret))

	o := Order{
		ordId:        100,
		customerName: "Jack",
		callback:     func() {},
	}
	InterfaceToStruct(o)
}
