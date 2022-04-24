// 三数之和
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 从小到大排序
	sort.Ints(nums)
	ans := make([][] int, 0)
	// 固定第一个，然后第二个第三个两边双指针遍历
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		third := len(nums)-1
		temp := -1 * nums[first]
		for second := first+1; second < third; second++{
			if (second > first + 1) && nums[second] == nums[second - 1]{
				continue
			}
			for second < third && nums[second] + nums[third] > temp {
				third--
			}
			if second == third {
				 break
			}
			if nums[second] + nums[third] == temp {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func main() {
	arr := []int{-1,0,1,2,-1,-4}
	res := threeSum(arr)
	fmt.Println(res)
}
