// 爬楼梯
package main

import (
	"fmt"
)

func climbStairs(n int) int {
	first, second, all := 0, 0 ,1
	for i := 1; i <= n; i++ {
		first = second
		second = all
		all = first + second
	}
	return all
}

func main() {
	res := climbStairs(2)
	fmt.Println(res)
}
