package main

import (
	"fmt"
	"strconv"
)

/*

题目描述
请实现如下接口

    /* 功能：四则运算

     * 输入：strExpression：字符串格式的算术表达式，如: "3+2*{1+2*[-4/(8-6)+7]}"

         * 返回：算术表达式的计算结果



public static int calculate(String strExpression)

{

请实现

return 0;

}

约束：

pucExpression字符串中的有效字符包括[‘0’-‘9’],‘+’,‘-’, ‘*’,‘/’ ,‘(’， ‘)’,‘[’, ‘]’,‘{’ ,‘}’。

pucExpression算术表达式的有效性由调用者保证;





输入描述:
输入一个算术表达式

输出描述:
得到计算结果

示例1
输入
复制
3+2*{1+2*[-4/(8-6)+7]}
输出
复制
25
*/

type stack []string

func (s *stack) push(v string) {
	*s = append(*s, v)
}

func (s *stack) top() string {
	return (*s)[len(*s)-1]
}

func (s *stack) pop() string {
	tail := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tail
}

func (s *stack) shift() string {
	h := (*s)[0]
	if len(*s) > 1 {
		*s = (*s)[1:]
	} else {
		*s = nil
	}
	return h
}

func (s *stack) length() int {
	return len(*s)
}

func (s *stack) is_empty() bool {
	return len(*s) == 0
}

type stack_int []int

func (s *stack_int) push(v int) {
	*s = append(*s, v)
}

func (s *stack_int) pop() int {
	tail := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tail
}

var op_rank = map[byte]int{
	'{': 1,
	'[': 1,
	'(': 1,
	'*': 3,
	'/': 3,
	'+': 2,
	'-': 2,
}

var bra_map = map[byte]byte{
	'}': '{',
	']': '[',
	')': '(',
}

func main() {
	var expr string
	//    for {
	_, err := fmt.Scan(&expr)
	if err != nil {
		return
	}
	s1 := stack{}
	s2 := stack{}
	last_is_op := false
	last_is_left := false
	for _, ch := range expr {
		ch_str := string(ch)
		if ch >= '0' && ch <= '9' {
			if !s1.is_empty() && !last_is_op {
				tmp := s1.pop()
				s1.push(tmp + ch_str)
			} else {
				s1.push(ch_str)
			}
			last_is_op = false
			last_is_left = false
			continue
		}

		if ch == '{' || ch == '[' || ch == '(' {
			s2.push(ch_str)
			last_is_op = true
			last_is_left = true
			continue
		}

		if ch == '-' && last_is_left {
			s1.push("0")
		}
		last_is_op = true
		last_is_left = false

		if ch == '}' || ch == ']' || ch == ')' {
			i := s2.pop()
			for {
				if s2.is_empty() || i[0] == bra_map[byte(ch)] {
					break
				}
				s1.push(i)
				i = s2.pop()
			}
			continue
		}

		if ch == '*' || ch == '/' || ch == '+' || ch == '-' {
			if !s2.is_empty() {
				i := s2.top()
				for op_rank[i[0]] >= op_rank[byte(ch)] {
					s1.push(i)
					s2.pop()
					if s2.is_empty() {
						break
					}
					i = s2.top()
				}
			}
			s2.push(ch_str)
		}
	}
	for !s2.is_empty() {
		s1.push(s2.pop())
	}
	re := stack_int{}
	for !s1.is_empty() {
		ch := s1.shift()
		if v, err := strconv.Atoi(ch); err == nil {
			re.push(v)
		} else {
			v2 := re.pop()
			v1 := re.pop()
			switch ch {
			case "+":
				re.push(v1 + v2)
			case "-":
				re.push(v1 - v2)
			case "*":
				re.push(v1 * v2)
			case "/":
				re.push(v1 / v2)
			}
		}
	}
	fmt.Println(re.pop())
	//    }
}
