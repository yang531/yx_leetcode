package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) []int {
	sort.Ints(nums)
	var back func(a []int, b int, c string)
	back = func(a []int, b int, c string) {
		
	}
	back(nums, 0, "")
	return []int{}
}



func main() {
	arr := []int{3, 2, 4}
	res := nextPermutation(arr)
	println(res)
}
