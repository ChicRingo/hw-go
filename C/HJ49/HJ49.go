package main

import (
	"fmt"
)

/*
题目描述
问题描述：有4个线程和1个公共的字符数组。
线程1的功能就是向数组输出A，线程2的功能就是向字符输出B，
线程3的功能就是向数组输出C，线程4的功能就是向数组输出D。
要求按顺序向数组赋值ABCDABCDABCD，ABCD的个数由线程函数1的参数指定。

输入描述:
输入一个int整数

输出描述:
输出多个ABCD

示例1
输入
复制
10
输出
复制
ABCDABCDABCDABCDABCDABCDABCDABCDABCDABCD
*/
func main1() {
	var n int
	for {
		num, _ := fmt.Scan(&n)
		if num == 0 {
			break
		}
		for i := 0; i < n; i++ {
			fmt.Printf("ABCD")
		}
		fmt.Printf("\n")
	}
}

var (
	str   []byte
	count int
	ch1   = make(chan int)
	ch2   = make(chan int)
	ch3   = make(chan int)
	ch4   = make(chan int)
	ch5   = make(chan int)
)

func main() {

	go addA()
	go addB()
	go addC()
	go addD()

	for {
		_, err := fmt.Scanf("%d\n", &count)
		if err != nil {
			break
		}

		ch1 <- 1
		<-ch5

		for i := range str {
			fmt.Printf("%c", str[i])
		}
		fmt.Println()
		str = nil
	}
}

func addA() {
	for {
		<-ch1
		if count != 0 {
			str = append(str, 'A')
			ch2 <- 1
		} else {
			ch5 <- 1
		}
		count--
	}
}

func addB() {
	for {
		<-ch2
		str = append(str, 'B')
		ch3 <- 1
	}
}

func addC() {
	for {
		<-ch3
		str = append(str, 'C')
		ch4 <- 1
	}
}

func addD() {
	for {
		<-ch4
		str = append(str, 'D')
		ch1 <- 1
	}
}
