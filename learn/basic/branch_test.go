package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	const filename = "abc.txt"
	t.Log("读取文件")
	//ReadFile输出两个值,一个是文件内容,一个是错误信息
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(contents))
		t.Log("将字段串根据字符分割成数组")
		strArray := strings.Split(string(contents), "\n\r")
		str := strings.Join(strArray, "|||||||||||||||||")
		fmt.Println(str)
		var arr []string
		t.Log("在切片末尾添加元素，并数组转成分隔符字符串")
		arrExplode := strings.Join(append(arr, "aa", "bb"), ",")
		fmt.Println(arrExplode)
	}
}

func TestAppend(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	a = append(a, 0) //先把原来的切片长度+1
	index := 2       //要把新元素插入到第二个位置
	copy(a[index+1:], a[index:])
	a[index] = 0 //新元素的值是0
	fmt.Println(a)
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		fmt.Printf("worong score:%d", score)
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}
	return g
}

func TestSwitch(t *testing.T) {
	fmt.Println(grade(0))
	fmt.Println(grade(101))
	fmt.Println(grade(1))
	fmt.Println(grade(90))
	fmt.Println(grade(80))
}
