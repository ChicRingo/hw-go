package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
题目描述
如果统计的个数相同，则按照ASCII码由小到大排序输出 。如果有其他字符，则对这些字符不用进行统计。

实现以下接口：
输入一个字符串，对字符中的各个英文字符，数字，空格进行统计（可反复调用）
按照统计个数由多到少输出统计结果，如果统计的个数相同，则按照ASCII码由小到大排序输出
清空目前的统计结果，重新统计
调用者会保证：
输入的字符串以‘\0’结尾。



输入描述:
输入一串字符。

输出描述:
对字符中的
各个英文字符（大小写分开统计），数字，空格进行统计，并按照统计个数由多到少输出,如果统计的个数相同，则按照ASII码由小到大排序输出 。如果有其他字符，则对这些字符不用进行统计。

示例1
输入
复制
aadddccddc
输出
复制
dca
*/
func main() {
	assicMax()
}
func assicMax() {
	r := bufio.NewReader(os.Stdin)
	for {

		btys, _, _ := r.ReadLine()
		if len(btys) == 0 {
			break
		}
		// 先将数据读入到此数组中
		arr := make([]int, 256)

		for _, b := range btys {
			arr[b]++
		}

		// 遍历数组，每次都找一个最大的值
		for {
			max := 0
			for j := 0; j < len(arr); j++ {
				if arr[j] == 0 {
					continue
				}
				if arr[j] > arr[max] {
					max = j
				}
			}
			if arr[max] == 0 {
				break
			}
			arr[max] = 0
			fmt.Printf("%c", max)
		}
		fmt.Println()
	}
}
