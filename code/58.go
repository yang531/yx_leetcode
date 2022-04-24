// 最后一个单词的长度
package main

import(
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")
	for i := len(arr) -1; i >=0; i--{
		//fmt.Println("content:", arr[i], "length:", len(arr[i]))
		if len(arr[i]) != 0 {
			return len(arr[i])
		}
	}
	return 0
}

func main() {
	var str = "   fly me   to   the moon  "
	res := lengthOfLastWord(str)
	fmt.Println(res)
}
