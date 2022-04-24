// 盛水最多的容器
package main

import (
	"fmt"
)

// 从前往后双循环计算会超时
// 选择从左右两端同步行走，哪边更小哪边前进一步
func maxArea(height []int) int {
	maxSum := 0
	left, right := 0, len(height)-1
	for left <= right{
		maxSum = maxNum(maxSum, minNum(height[left], height[right]) * (right-left))
		if height[left] < height[right] {
			left++
		}else {
			right--
		}
	}

	return maxSum
}

func maxNum(a int, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

func minNum(a int, b int) int {
	if a < b {
		return a
	}else {
		return b
	}
}

func main() {
	height := []int{1,1}
	res := maxArea(height)
	fmt.Println(res)
}
