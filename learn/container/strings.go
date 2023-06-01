package main

import (
	"fmt"
	"unicode/utf8"
)
//rune相当于go的char
//使用 utf8.RuneCountInString获得字符数量
//使用len获得字节长度
//使用[]byte获得字节
func main() {
	s := "Yes我爱你赛事"
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %x)", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune connt:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	fmt.Println(len(bytes))
	fmt.Println(len(s))
	fmt.Println(bytes)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()
	for i,ch  := range []rune(s){
		fmt.Printf("(%d %c)",i,ch)
	}
	fmt.Println()
}
