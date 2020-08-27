package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make([]int, 5)
	b := a[:3]
	for _, v := range b {
		v = v + 1
	}
	fmt.Println(a)
	fmt.Println(b)
	for i, v := range b {
		a[i] = v + 1
	}
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(strings.Index("qwerasdf", "asdf"))
	fmt.Println(strStr("qwerasdf", "asdf"))
}
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	// i不需要到len-1
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 判断字符串长度是否相等
		if len(needle) == j {
			return i
		}
	}
	return -1
}
