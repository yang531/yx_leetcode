// 移除元素
package main

import (
	"fmt"
)


func removeElement(nums []int, val int) int {
	left := 0
	for _, value := range nums{
		if value != val {
			// 题目要求需要原地修改数组
			 nums[left] = value
			 left++
		}
	}
	fmt.Println(nums)
	return left
}


func main()  {
	nums := [] int{3, 2, 2, 3}
	res := removeElement(nums, 3)
	fmt.Println(res)
}

