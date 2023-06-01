package main

import (
	"fmt"
	"testing"
)

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		fmt.Println(i, ch)
		lastOccurred[ch] = i
	}
	return maxLength
}

func TestNon(t *testing.T) {
	fmt.Println(lengthOfNonRepeatingSubStr("qwerrrwewewwertyuiop"))
}
