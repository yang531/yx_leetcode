// 搜索插入位置
package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	for index, value := range nums{
		if value >= target {
			return index
		}
	}
	return len(nums)
}

func main() {
	nums := [] int{1, 2, 3, 8}
	res := searchInsert(nums, 6)
	fmt.Println(res)
}
