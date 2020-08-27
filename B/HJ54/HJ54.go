package main

import "fmt"

/*
题目描述
给定一个字符串描述的算术表达式，计算出结果值。

输入字符串长度不超过100，合法的字符包括”+, -, *, /, (, )”，”0-9”，字符串内容的合法性及表达式语法的合法性由做题者检查。本题目只涉及整型计算。

输入描述:
输入算术表达式

输出描述:
计算出结果值

示例1
输入
复制
400+5
输出
复制
405
*/

func main() {
	for {
		str := ""
		n, _ := fmt.Scan(&str)
		if n == 0 {
			break
		}
		weight := make(map[rune]int)

		weight['+'] = 1
		weight['-'] = 1
		weight['*'] = 2
		weight['/'] = 2

		arr := []rune(str)
		data := make([]int, 0)
		operator := make([]rune, 0)

		for i := 0; i < len(arr); i++ {
			if arr[i] >= '0' && arr[i] <= '9' {
				digit := 0
				for i < len(arr) && arr[i] >= '0' && arr[i] <= '9' {
					digit = digit*10 + int(arr[i]-'0')
					i++
				}
				i--
				data = append(data, digit)
			} else {
				if arr[i] == '-' {
					if i-1 < 0 || arr[i-1] == '{' || arr[i-1] == '[' || arr[i-1] == '(' {
						data = append(data, 0)
					}
				}
				if len(operator) == 0 {
					operator = append(operator, arr[i])
				} else if (weight[operator[len(operator)-1]] < weight[arr[i]]) || arr[i] == '{' || arr[i] == '[' || arr[i] == '(' {
					operator = append(operator, arr[i])
				} else {
					if arr[i] == '}' || arr[i] == ']' || arr[i] == ')' {
						for operator[len(operator)-1] != '{' && operator[len(operator)-1] != '[' && operator[len(operator)-1] != '(' {
							op := operator[len(operator)-1]
							operator = operator[:len(operator)-1]
							data1 := data[len(data)-2]
							data2 := data[len(data)-1]
							data = data[:len(data)-2]

							res := compute(data1, data2, op)
							data = append(data, res)
						}
						operator = operator[:len(operator)-1]
					} else {
						for len(operator) > 0 && weight[operator[len(operator)-1]] >= weight[arr[i]] {
							op := operator[len(operator)-1]
							operator = operator[:len(operator)-1]
							data1 := data[len(data)-2]
							data2 := data[len(data)-1]
							data = data[:len(data)-2]
							res := compute(data1, data2, op)
							data = append(data, res)
						}
						operator = append(operator, arr[i])
					}
				}
			}
		}
		for len(operator) > 0 {
			op := operator[len(operator)-1]
			operator = operator[:len(operator)-1]
			data1 := data[len(data)-2]
			data2 := data[len(data)-1]
			data = data[:len(data)-2]
			res := compute(data1, data2, op)
			data = append(data, res)
		}
		fmt.Println(data[0])
	}
}

func compute(data1, data2 int, op rune) int {
	res := 0
	switch op {
	case '+':
		res = data1 + data2
		break
	case '-':
		res = data1 - data2
		break
	case '*':
		res = data1 * data2
		break
	case '/':
		res = data1 / data2
		break
	}
	return res
}
