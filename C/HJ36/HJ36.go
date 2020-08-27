package main

import (
	"fmt"
	"strings"
)

/*
题目描述
有一种技巧可以对数据进行加密，它使用一个单词作为它的密匙。下面是它的工作原理：首先，选择一个单词作为密匙，如TRAILBLAZERS。如果单词中包含有重复的字母，只保留第1个，其余几个丢弃。现在，修改过的那个单词属于字母表的下面，如下所示：

A B C D E F G H I J K L M N O P Q R S T U V W X Y Z

T R A I L B Z E S C D F G H J K M N O P Q U V W X Y

上面其他用字母表中剩余的字母填充完整。在对信息进行加密时，信息中的每个字母被固定于顶上那行，并用下面那行的对应字母一一取代原文的字母(字母字符的大小写状态应该保留)。因此，使用这个密匙，Attack AT DAWN(黎明时攻击)就会被加密为Tpptad TP ITVH。

请实现下述接口，通过指定的密匙和明文得到密文。

输入描述:
先输入key和要加密的字符串

输出描述:
返回加密后的字符串

示例1
输入
复制
nihao
ni
输出
复制
le

*/
func main() {
	var key string
	var data string
	for {
		_, err := fmt.Scan(&key)
		if err != nil {
			break
		}
		_, err = fmt.Scan(&data)
		if err != nil {
			break
		}

		test(key, data)
	}
}
func test(key, data string) {
	keyUp := strings.ToUpper(key)
	dataUp := strings.ToUpper(data)
	fmt.Println(keyUp, dataUp)

	tmp := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	k := []byte(keyUp)
	//da := []rune(data)
	fmt.Println(string(tmp))
	fmt.Println(len(keyUp))
	fmt.Println(k)
	keyTmp := ""
	// 遍历tmp和keyUp，将tmp中和keyUp重复元素剔除
	for _, v1 := range tmp {
	b:
		for _, v2 := range k {
			if v1 != v2 {
				keyTmp += string(v2)
				break b
			}
		}
	}
	fmt.Println(keyTmp)

	//

	//
	char := make(map[byte]bool, len(keyUp))
	for _, v := range k {
		up := true
		if v >= 'a' && v <= 'z' {
			up = false
		}
		char[v] = up
		fmt.Println(string(v))
	}
	//sort.Strings(char)

}
