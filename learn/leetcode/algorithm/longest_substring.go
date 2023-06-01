package algorithm

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	arr := strings.Split(s, "")
	start := 0
	tmpSubstringArr := make(map[string]int)
	longestStringMax := 0
	tmpLongestStringMax := 0
	for k, v := range arr {
		if val, ok := tmpSubstringArr[v]; ok && val >= start {
			tmpLongestStringMax = k - start

			if tmpSubstringArr[v]+1 > start {
				start = tmpSubstringArr[v] + 1
			}
		} else {
			tmpLongestStringMax = k - start + 1
		}

		if tmpLongestStringMax > longestStringMax {
			longestStringMax = tmpLongestStringMax
		}
		
		tmpSubstringArr[v] = k
	}
	return longestStringMax
}

func lengthOfLongestSubstringV2(s string) int {
	r, l := 0, 0
	var ret int
	for i := range s {
		index := strings.Index(s[l:i], string(s[i]))
		if index == -1 {
			r++
		} else {
			r = i + 1
			l += index + 1
		}
		ret = max(len(s[l:r]), ret)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
