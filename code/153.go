// 寻找旋转排序数组中的最小值
package main

import (
	"fmt"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left < right {
		mid = (left + right) >>1
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func Intmin(a int, b int)  int{
	if a > b {
		return b
	} else{
		return a
	}
}

func main() {
	res := findMin([]int{3,1,2})
	fmt.Println(res)
}
