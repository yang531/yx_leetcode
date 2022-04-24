package main

import (
	"fmt"
)


func longestCommonPrefix(strs []string) string {
	val := strs[0]
	for i := 1; i < len(strs); i++{
		length := min(len(strs[i]), len(val))
		temp := ""
		for j := 0; j < length; j++{
			if strs[i][j] == val[j]{
				 temp = temp + string(strs[i][j])
			} else {
				break
			}
		}
		val = temp
		length = len(val)
	}
	return val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main()  {
	//strs := [] string{"dog", "racecar", "car"}
	res := longestCommonPrefix([] string{"dog","racecar","car"})

	fmt.Println(res)
}

