package main

import (
	"fmt"
	"test_go_mod/learn/interface/testing"
)

//接口的概念以及简单demo
//定义一个接口，包含一个函数Get
//接口是系列方法的集合
//单个类型可以实现多个方法
//同一个接口可以被多种类型实现
type retriever interface {
	Get(string) string
}


func getRetriever() retriever {
	return testing.Retriever{}
	//return infra.Retriever{}
}

func main() {
	retriever := getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com"))
}
