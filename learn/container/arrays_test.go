package main

import (
	"fmt"
	"testing"
)

func TestArrays(t *testing.T) {
	t.Log("定义数组")
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)
	t.Log("遍历数组")
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 {
		fmt.Println(i, v)
		fmt.Println(i, v)
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}
	printArray(&arr3)
	fmt.Println(arr3)
}

//数组是值类型,加指针
func printArray(arr *[5]int) {
	arr[0] = 100
	var arr3 [5]int
	arr3[1] = 11
	arr3[2] = 110
	fmt.Println(arr3)
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
