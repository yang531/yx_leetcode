// 搜索旋转排序数组
/*
要求时间复杂度为O(n)，使用二分查找
寻找中间值后判断前半部分和后半部分哪边是有序的，然后再使用二分法
*/
package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := -1
	for left <= right {
		mid = (left+right+1) >>1
		if nums[mid] == target {
			return mid
		}
		// 前半部分有序
		if nums[left] < nums[mid]{
			if nums[left] <= target && nums[mid] > target  {
				right = mid - 1
			}else {
				left = mid + 1
			}
		} else { // 后半部分有序
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			}else {
				right = mid - 1
			}
		}
	}
	return -1
}


func main() {
	nums := [] int{3, 4, 5, 6, 1, 2}
	res := search(nums, 7)
	fmt.Println(res)
}
