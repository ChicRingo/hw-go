package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		test(input.Text())
	}
	//reader := bufio.NewReader(os.Stdin)
	//for {
	//	line, err := reader.ReadString('\n')
	//	if err != nil {
	//		return
	//	}
	//	test(line)
	//}
}

func CheckValidStep(str string) (int, int) {
	if len(str) >= 4 || len(str) <= 1 {
		return 0, 0
	}
	step, _ := strconv.Atoi(str[1:])
	fmt.Println(step)
	if str[0] == 'A' {
		return 1, -step
	} else if str[0] == 'D' {
		return 1, step
	} else if str[0] == 'W' {
		return 2, step
	} else if str[0] == 'S' {
		return 2, -step
	}
	return 0, 0
}

func test(line string) {
	lines := strings.Split(line, ";")
	x := 0
	y := 0
	for _, line := range lines {
		t, step := CheckValidStep(line)
		if 1 == t {
			x += step
		} else if 2 == t {
			y += step
		}
	}
	//fmt.Println(x, ',', y)
	fmt.Printf("%d%c%d\n", x, ',', y)
}
