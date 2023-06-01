package main

import "fmt"

//defer有一个栈、先进后出
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	return
	fmt.Println(4)
}




func main() {
	tryDefer()
}
