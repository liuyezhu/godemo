package main

import (
	"fmt"
	"reflect"
)

type CustomValidateFunc struct {
}

func main() {
	var validateFunc CustomValidateFunc
	vf := reflect.ValueOf(&validateFunc)
	fmt.Println(vf.MethodByName("FuncA").Call(valOf(1, "AAAA")))
	fmt.Println(vf.MethodByName("FuncB").Call(valOf(2, "BBBB")))

}

func (*CustomValidateFunc) FuncA(a int, b string) string {
	return fmt.Sprintf("成功调用FuncA:参数a是%d,参数b是%s", a, b)
}

func (*CustomValidateFunc) FuncB(a int, b string) string {
	return fmt.Sprintf("成功调用FuncB:参数a是%d,参数b是%s", a, b)
}

func valOf(i ...interface{}) []reflect.Value {
	var rt []reflect.Value
	for _, i2 := range i {
		rt = append(rt, reflect.ValueOf(i2))
	}
	return rt
}
