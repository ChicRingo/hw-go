package main

import (
	"fmt"
	"strings"
)

/*
题目描述
问题描述：在计算机中，通配符一种特殊语法，广泛应用于文件搜索、数据库、正则表达式等领域。现要求各位实现字符串通配符的算法。
要求：
实现如下2个通配符：
*：匹配0个或以上的字符（字符由英文字母和数字0-9组成，不区分大小写。下同）
？：匹配1个字符


输入：
通配符表达式；
一组字符串。


输出：
返回匹配的结果，正确输出true，错误输出false

输入描述:
先输入一个带有通配符的字符串，再输入一个需要匹配的字符串

输出描述:
返回匹配的结果，正确输出true，错误输出false

示例1
输入
复制
te?t*.*
txt12.xls
输出
复制
false
*/
func main() {
	for {
		var p, s string
		n, _ := fmt.Scanln(&p)
		if n == 0 {
			break
		}
		fmt.Scanln(&s)
		fmt.Println(getResult(p, s))
	}
}

func getResult(p, s string) bool {
	if s == "" {
		p = strings.ReplaceAll(p, "*", "")
		if len(p) == 0 {
			return true
		}
		return false
	}
	if p == "" && s != "" {
		return false
	}
	if p[0] == '*' {
		return getResult(p[1:], s) || getResult(p, s[1:])
	} else if p[0] == '?' {
		return getResult(p[1:], s[1:])
	} else {
		if p[0] != s[0] {
			return false
		}
		return getResult(p[1:], s[1:])
	}
}
