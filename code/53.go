package main

import(
	"fmt"
)

func maxSubArray(nums []int) int {
	var maxNum = nums[0]
	for i := 1; i < len(nums); i++{
		// 从前往后挨个计算，保留前面计算过的最大值
		if nums[i] + nums[i-1] > nums[i]{
			nums[i] += nums[i-1]
		}
		if maxNum < nums[i]{
			maxNum = nums[i]
		}
	}
	return maxNum
}


func main() {
	num := []int{-2, 3, 3, 4}
	res := maxSubArray(num)
	fmt.Println(res)
}
