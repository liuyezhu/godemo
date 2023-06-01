package main

import (
	"fmt"
	"testing"
)

//函数外层不能使用：定义变量
var (
	aa = 3
	ss = "kkk"
	cc = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b, c int = 1, 2, 3
	var s string = "abc"
	fmt.Println(a, b, c, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 2, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//推荐使用
func variableShorter() {
	a, b, c, s := 2, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func TestVariable(t *testing.T) {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, cc)
}
