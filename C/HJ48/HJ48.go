package main

import "fmt"

/*
题目描述
输入一个单向链表和一个节点的值，从单向链表中删除等于该值的节点，删除后如果链表中无节点则返回空指针。

链表结点定义如下：

struct ListNode

{

int       m_nKey;

ListNode* m_pNext;

};

详细描述：

本题为考察链表的插入和删除知识。

链表的值不能重复

构造过程，例如

1 <- 2

3 <- 2

5 <- 1

4 <- 5

7 <- 2

最后的链表的顺序为 2 7 3 1 5 4

删除 结点 2

则结果为 7 3 1 5 4



输入描述:
1 输入链表结点个数
2 输入头结点的值
3 按照格式插入各个结点
4 输入要删除的结点的值

输出描述:
输出删除结点后的序列，每个数后都要加空格

示例1
输入
复制
5
2
3 2
4 3
5 2
1 4
3
输出
复制
2 5 4 1
*/
type List struct {
	value int
	next  *List
}

func main() {
	for {
		count := 0
		n, _ := fmt.Scan(&count)
		if n == 0 {
			break
		}
		headNode := 0
		fmt.Scan(&headNode)
		node1, node2 := 0, 0
		head := &List{
			value: headNode,
		}

		m := make(map[int]*List)
		m[headNode] = head
		for {
			fmt.Scan(&node1, &node2)
			if v, ok := m[node2]; ok {
				nextNode := v.next
				newNode := &List{
					value: node1,
					next:  nextNode,
				}
				v.next = newNode
				m[node1] = newNode
			}
			if len(m) == count {
				break
			}
		}
		delNode := 0
		fmt.Scan(&delNode)
		cur := head
		for cur != nil {
			if cur.value != delNode {
				fmt.Printf("%d ", cur.value)
			}
			cur = cur.next
		}
		fmt.Println()
	}
}
