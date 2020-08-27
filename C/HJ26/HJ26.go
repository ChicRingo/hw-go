package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
题目描述
编写一个程序，将输入字符串中的字符按如下规则排序。

规则 1 ：英文字母从 A 到 Z 排列，不区分大小写。

如，输入： Type 输出： epTy

规则 2 ：同一个英文字母的大小写同时存在时，按照输入顺序排列。

如，输入： BabA 输出： aABb

规则 3 ：非英文字母的其它字符保持原来的位置。


如，输入： By?e 输出： Be?y


注意有多组测试数据，即输入有多行，每一行单独处理（换行符隔开的表示不同行）


输入描述:
输入字符串
输出描述:
输出字符串
示例1
输入
复制
A Famous Saying: Much Ado About Nothing (2012/8).
输出
复制
A aaAAbc dFgghh: iimM nNn oooos Sttuuuy (2012/8).
*/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		sort(input.Text())
	}
}

func sort(s string) {
	symbolIndex := make([]int, 0)
	symbolValue := make([]byte, 0)
	//遍历字符判断，非字母就存index和value到两个数组里
	for i, v := range s {
		if v > 'z' || v > 'Z' && v < 'a' || v < 'A' {
			symbolIndex = append(symbolIndex, i)
			symbolValue = append(symbolValue, uint8(v))
		}
	}
	//fmt.Println(symbolIndex, symbolValue)
	// 遍历26个大小写字母，按a-z顺序存入数组char中
	char := make([]byte, 0)
	for i := 0; i < 26; i++ {
		for _, v := range s {
			value := uint8(v)
			if value == 'a'+uint8(i) || value == 'A'+uint8(i) {
				char = append(char, value)
			}
		}
	}
	//fmt.Println(char)
	//遍历符号位置信息存入结果数组中
	res := make([]byte, len(s))
	for i, v := range symbolValue {
		res[symbolIndex[i]] = v
	}
	//fmt.Println(res)
	//遍历数组若元素为0 则赋值为字符byte
	for _, c := range char {
		for i, v := range res {
			if v == 0 {
				res[i] = c
				break
			}
		}
	}
	fmt.Println(string(res))
}
