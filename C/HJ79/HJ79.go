package main

import "fmt"

/*
题目描述
对于不同的字符串，我们希望能有办法判断相似程度，我们定义了一套操作方法来把两个不相同的字符串变得相同，具体的操作方法如下：

1 修改一个字符，如把“a”替换为“b”。

2 增加一个字符，如把“abdd”变为“aebdd”。

3 删除一个字符，如把“travelling”变为“traveling”。

比如，对于“abcdefg”和“abcdef”两个字符串来说，我们认为可以通过增加和减少一个“g”的方式来达到目的。上面的两种方案，都只需要一次操作。把这个操作所需要的次数定义为两个字符串的距离，而相似度等于“距离＋1”的倒数。也就是说，“abcdefg”和“abcdef”的距离为1，相似度为1/2=0.5.

给定任意两个字符串，你是否能写出一个算法来计算出它们的相似度呢？



请实现如下接口

 /* 功能：计算字符串的相似度
  * 输入：pucAExpression/ pucBExpression：字符串格式，如: "abcdef"
  * 返回：字符串的相似度,相似度等于“距离＋1”的倒数,结果请用1/字符串的形式,如1/2



输入描述:
输入两个字符串

输出描述:
输出相似度，string类型

示例1
输入
复制
abcdef
abcdefg
输出
复制
1/2
*/
func main() {
	for {
		var s1, s2 string
		n, _ := fmt.Scanln(&s1)
		if n == 0 {
			break
		}
		fmt.Scanln(&s2)
		l1 := len(s1)
		l2 := len(s2)
		arr := make([][]int, l1+1)
		cost := 0
		for i := 0; i <= l1; i++ {
			arr[i] = make([]int, l2+1)
		}
		for i := 0; i <= l1; i++ {
			arr[i][0] = i
		}
		for i := 0; i <= l2; i++ {
			arr[0][i] = i
		}
		for i := 1; i <= l1; i++ {
			for j := 1; j <= l2; j++ {
				if s1[i-1] == s2[j-1] {
					cost = 0
				} else {
					cost = 1
				}
				arr[i][j] = min(arr[i][j-1]+1, arr[i-1][j]+1, arr[i-1][j-1]+cost)
			}
		}
		fmt.Printf("%d/%d\n", 1, (arr[l1][l2] + 1))
	}
}
func min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}
