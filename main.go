package main

import (
	"fmt"
	"sync"
)

// 使用2个goroutine轮流打印一个数组
var wg sync.WaitGroup
var ch = make(chan bool)
var send = make(chan int)

type S struct {
	string
}

func init() {
	fmt.Println(1)

}
func main() {
	fmt.Println("2")
}
func main2() {
	a := ""
	bs := []byte(a)
	fmt.Println(len(bs), cap(bs))
}
func main1() {

	wg.Add(3)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := len(nums)
	go Send(nums, l) // 发送
	go printOdd(l)   // 打印奇数
	go printEven(l)  // 打印偶数
	wg.Wait()
}

// 发送数组元素
func Send(nums []int, l int) {
	defer wg.Done()
	for i := 0; i < l; i++ {
		a := nums[i]
		send <- a
	}
}

// 打印奇数
func printEven(l int) {
	defer wg.Done()
	for i := 0; i < l; i++ {
		ch <- true
		if i&1 == 0 {
			a := <-send
			fmt.Println("printEven:", a)
		}
	}
}

// 打印偶数
func printOdd(l int) {
	defer wg.Done()
	for i := 0; i < l; i++ {
		<-ch
		if i&1 == 1 {
			a := <-send
			fmt.Println("printOdd: ", a)
		}
	}
}

//// 递归算法，看以 Node 为根的二叉查找树是否包含e
//func contains(n *Node, e interface{}) bool {
//	if n == nil {
//		return false
//	}
//
//	if util.Compare(n.e, e) == 0 {
//		return true
//	} else if util.Compare(e, n.e) < 0 {
//		return contains(n.left, e)
//	} else {
//		return contains(n.right, e)
//	}
//}
//
//// 递归算法，插入元素e，返回插入新节点之后的二叉查找树的根
//func (b *BST) add(n *Node, e interface{}) *Node {
//	if n == nil {
//		b.size++
//		return createNode(e)
//	}
//
//	if util.Compare(e, n.e) < 0 {
//		n.left = b.add(n.left, e)
//	} else if util.Compare(e, n.e) > 0 {
//		n.right = b.add(n.right, e)
//	}
//	return n
//}
