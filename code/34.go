package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] != target {
			left ++
		}
		if nums[right] != target {
			right--
		}
		if nums[left] == nums[right] && nums[left] == target {
			break
		}
	}
	if nums[left] != target{
		left, right = -1, -1
	}
	return []int{left, right}
}

func searchRange1(nums []int, target int) []int {
	// SearchInts函数索引数组中的target的位置，数组需要是升序排列的
	// 如果数组中没有目标值，则返回遇见的第一个比目标大的值的位置
	left := sort.SearchInts(nums, target)
	// left == len(nums) 避免数组为空的情况下报错
	if left == len(nums) || nums[left] != target {
		return [] int{-1, -1}
	}
	right := sort.SearchInts(nums, target+1)-1
	return []int{left, right}
}



func main() {
	res := searchRange1([]int{1, 3, 2, 3}, 1)
	fmt.Println(res)
}
