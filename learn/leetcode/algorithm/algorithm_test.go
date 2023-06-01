package algorithm

import (
	"fmt"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	maxLen := lengthOfLongestSubstring("aabaab!bb")
	//maxLen := lengthOfLongestSubstring("pwwkew")
	fmt.Println(maxLen)
}
