package main

import "fmt"

/*
题目描述
MP3 Player因为屏幕较小，显示歌曲列表的时候每屏只能显示几首歌曲，用户要通过上下键才能浏览所有的歌曲。为了简化处理，假设每屏只能显示4首歌曲，光标初始的位置为第1首歌。



现在要实现通过上下键控制光标移动来浏览歌曲列表，控制逻辑如下：

歌曲总数<=4的时候，不需要翻页，只是挪动光标位置。

光标在第一首歌曲上时，按Up键光标挪到最后一首歌曲；光标在最后一首歌曲时，按Down键光标挪到第一首歌曲。



其他情况下用户按Up键，光标挪到上一首歌曲；用户按Down键，光标挪到下一首歌曲。



  2. 歌曲总数大于4的时候（以一共有10首歌为例）：



特殊翻页：屏幕显示的是第一页（即显示第1 – 4首）时，光标在第一首歌曲上，用户按Up键后，屏幕要显示最后一页（即显示第7-10首歌），同时光标放到最后一首歌上。同样的，屏幕显示最后一页时，光标在最后一首歌曲上，用户按Down键，屏幕要显示第一页，光标挪到第一首歌上。



一般翻页：屏幕显示的不是第一页时，光标在当前屏幕显示的第一首歌曲时，用户按Up键后，屏幕从当前歌曲的上一首开始显示，光标也挪到上一首歌曲。光标当前屏幕的最后一首歌时的Down键处理也类似。



其他情况，不用翻页，只是挪动光标就行。



输入描述:
输入说明：
1 输入歌曲数量
2 输入命令 U或者D

输出描述:
输出说明
1 输出当前列表
2 输出当前选中歌曲

示例1
输入
复制
10
UUUU
输出
复制
7 8 9 10
7
*/
const page = 4

func main() {
	var (
		n      int    // 歌曲总数
		action string // 操作
	)
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		fmt.Scan(&action)

		cursor := 0 // 当前指向的歌曲,从 0 开始计数
		if n <= page {
			for _, ch := range action {
				if ch == 'U' {
					cursor = (cursor + 1) % n
				} else {
					cursor = (cursor - 1) % n
				}
			}
			for i := 0; i < n; i++ {
				fmt.Printf("%d ", i+1)
			}
			fmt.Println()
			fmt.Println(cursor + 1)
			continue
		}

		header := 0 // 当前列表的头
		for _, ch := range action {
			if ch == 'U' {
				if cursor == 0 {
					header = n - page
					cursor = n - 1
				} else if cursor == header {
					// 光标在当前页的第一首
					cursor--
					header--
				} else {
					cursor--
				}
			} else {
				if cursor == n-1 {
					header = 0
					cursor = 0
				} else if cursor == header+page-1 {
					// 当前光标在当前页的最后一首
					cursor++
					header++
				} else {
					cursor++
				}
			}
		}

		for i := 1; i <= page; i++ {
			fmt.Printf("%d ", header+i)
		}
		fmt.Println()
		fmt.Println(cursor + 1)
	}
}
