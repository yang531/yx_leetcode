// 四数之和
package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	arr := make([][]int, 0)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first+1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			third, fourth := second+1, len(nums) - 1
			for third < fourth{
				temp := nums[first] + nums[second] + nums[third] + nums[fourth]
				if temp == target {
					arr = append(arr, []int{nums[first], nums[second], nums[third], nums[fourth]})
					//third++
					//fourth--
					for third ++; third < fourth && nums[third] == nums[third-1]; third++ {
					}
					for fourth --; third < fourth && nums[fourth] == nums[fourth+1]; fourth-- {
					}
				}else if temp < target {
					third++
				}else {
					fourth--
				}

			}
		}
	}
	return arr
}

func main() {
	arr := [] int{-2,-1,-1,1,1,2,2}
	res := fourSum(arr, 0)
	fmt.Println(res)
}
