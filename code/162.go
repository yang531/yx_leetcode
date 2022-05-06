// 寻找峰值
/* 二分查找法寻找数组中的峰值（峰值指其值严格大于左右的元素）
*/

package main

import (
	"fmt"
	"math"
)

func findPeakElement(nums []int) int {
	left, right := 0, len(nums) -1

	get := func(index int) int{
		if index >= len(nums) || index < 0 {
			return math.MinInt64
		} else {
			return nums[index]
		}
	}
	for{
		mid := (left + right) >> 1
		if get(mid - 1) < get(mid) && get(mid) > get(mid + 1) {
			return mid
		}
		if get(mid) < get(mid + 1) {
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
}

func main() {
	res := findPeakElement([] int{1,2})
	fmt.Println(res)
}
