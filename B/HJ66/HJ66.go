package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands = [6][2]string{
	{"reset"},
	{"reset", "board"},
	{"board", "add"},
	{"board", "delete"},
	{"reboot", "backplane"},
	{"backplane", "abort"},
}

// 命令的长度
var cLength = [...]int{1, 2, 2, 2, 2, 2}

var actions = [...]string{
	"reset what",
	"board fault",
	"where to add",
	"no board at all",
	"impossible",
	"install first",
}

const unknown = "unknown command"

func match(str []string) int {
	l := len(str)
	index := -1 // 搜索到命令的索引值

	for ix, val := range commands {
		if l != cLength[ix] {
			continue
		}
		if l == 1 && strings.Contains(val[0], str[0]) {
			if index == -1 {
				index = ix
			} else {
				index = -1
				break
			}
		} else if l == 2 && strings.Contains(val[0], str[0]) && strings.Contains(val[1], str[1]) {
			if index == -1 {
				index = ix
			} else {
				index = -1
				break
			}
		}

	}
	return index
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, _ := reader.ReadLine()
		if len(input) == 0 {
			break
		}
		str := string(input)
		i := match(strings.Fields(str))
		if i == -1 {
			fmt.Println(unknown)
		} else {
			fmt.Println(actions[i])
		}
	}
}
