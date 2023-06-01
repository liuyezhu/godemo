package main

/**
动态函数调用
不能够完美支持golang的多返回值。
只能通过 for range []reflect.value 来获得多返回值，
而不是可以 ret1, ret2 := Apply(f, args) 来接收多返回值
*/

import (
	"fmt"
	"reflect"
)

type CustomValidateFunc struct {
}

func main() {
	var validateFunc CustomValidateFunc
	vf := reflect.ValueOf(&validateFunc)
	fmt.Println(vf)
	vft := vf.Type()
	//读取方法数量
	//mNum := vf.NumMethod()
	//fmt.Println(mNum)
	fmt.Println(1212415)
	fmt.Println(vft.Method(0))
	//fmt.Println(454545)
	fmt.Println(vf.MethodByName("Check").Call(valOf(1, "成功调用Check方法")))
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	//ret1 := Apply(Cal, []interface{}{2, "sdsd"})
	//for _, v := range ret1 {
	//	fmt.Println(v)
	//}
	//fmt.Println(ret1)
}

func (*CustomValidateFunc) Check(a int, b string) string {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("我是check")
	return "ssed"
}

func valOf(i ...interface{}) []reflect.Value {
	var rt []reflect.Value
	for _, i2 := range i {
		rt = append(rt, reflect.ValueOf(i2))
	}
	return rt
}

//func Apply(f interface{}, args []interface{}) []reflect.Value {
//	fun := reflect.ValueOf(f)
//	in := make([]reflect.Value, len(args))
//	for k, param := range args {
//		in[k] = reflect.ValueOf(param)
//	}
//	r := fun.Call(in)
//	return r
//}

//func Cal(a int, b int) (int, string) {
//	return a + b, "hello"
//}
