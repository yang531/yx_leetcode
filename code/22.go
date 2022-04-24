package main

import (
	"fmt"
)



func generateParenthesis(n int) []string {
	var arr []string
	var dfs func(left, right int, item string)
	dfs = func(left, right int, item string) {
		if left == right && left == n {
			arr = append(arr, item)
			return
		}
		if right > left || left > n || right > n {
			return
		}
		dfs(left+1, right, item+"(")
	}

	dfs(0, 0, "")
	return arr
}

func main() {
	res := generateParenthesis(4)
	fmt.Println(res)
}
