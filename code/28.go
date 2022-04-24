package main

import (
	"fmt"
)

func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	flag:
	for i := 0; i+m <= n; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				// continue 跳出循环至最外层
				continue flag
			}
		}
		return i
	}
	return -1
}

func main() {
	res := strStr("aaaab", "")
	fmt.Println(res)
}
