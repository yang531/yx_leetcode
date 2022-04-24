// 下一个排列
package main

import (
	"sort"
)

func nextPermutation(nums []int) []int {
	sort.Ints(nums)

	return []int{}
}



func main() {
	arr := []int{3, 2, 4}
	res := nextPermutation(arr)
	println(res)
}
