package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{1, 2}
	a = append([]int{1}, a...)
	fmt.Println(a)
	var arr []string
	arr = append(arr, "a")
	arr = append([]string{"aa"}, arr...)
	fmt.Println(arr)
	t.Log("向切片末尾添加元素")
	arr = append(arr, "2222")
	t.Log("向切片任意位置添加元素")
	index := 1 //要把新元素插入到第二个位置
	copy(arr[index+1:], arr[index:])
	arr[index] = "任意位置"
	fmt.Println(arr)
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := array[2:6]
	s2 := s1[3:5]
	fmt.Println(s2[1])
	fmt.Println(s1, s2)
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))
	//s2 = append(s2, 1, 2, 3, 4, 7, 7, 7)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5", s3, s4, s5)
	fmt.Println("arr =", array)

}

func TestSlice1(t *testing.T) {
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	s1 = append(s1, 10)
	printSlice(s1)
	s2 := make([]int, 16)
	s3 := make([]int, 16, 32)
	printSlice(s2)
	printSlice(s3)
	t.Log("复制slice")
	copy(s2, s1)
	printSlice(s2)
	t.Log("删除指定位置slice元素")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
	t.Log("删除头部slice元素")
	front := s2[0]
	s2 = s2[1:]
	printSlice(s2)
	fmt.Println(front)
	t.Log("删除尾巴slice元素")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))
}
