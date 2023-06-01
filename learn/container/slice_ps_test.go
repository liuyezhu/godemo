package main

import (
	"fmt"
	"testing"
)

func TestSlicePs(t *testing.T) {
	var s []int //Zero  value for slice is nil
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

}
