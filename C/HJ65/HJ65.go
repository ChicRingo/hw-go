package main

import (
	"fmt"
	"strings"
)

/*
题目描述
查找两个字符串a,b中的最长公共子串。若有多个，输出在较短串中最先出现的那个。
输入描述:
输入两个字符串
输出描述:
返回重复出现的字符
示例1
输入
复制
abcdefghijklmnop
abcsafjklmnopqrstuvw
输出
复制
jklmnop
*/
func main() {
	for {
		var (
			s1 string
			s2 string
		)
		_, err := fmt.Scan(&s1, &s2)
		if err != nil {
			break
		}
		if len(s1) > len(s2) {
			tmp := s2
			s2 = s1
			s1 = tmp
		}
		s3 := ""
		max := 0
		for i := 0; i < len(s1)-1; i++ {
			for j := i + 1; j <= len(s1); j++ {
				if strings.Contains(s2, s1[i:j]) && max < j-i {
					max = j - i
					s3 = s1[i:j]
				}
			}

		}
		fmt.Printf("%s\n", s3)
	}
}
