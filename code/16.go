// 最接近的三数之和
package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	aim := math.MaxInt32
	update := func(temp int) {
		if abs(aim - target) > abs(temp - target) {
			aim = temp
		}
	}
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		third := len(nums) - 1
		for second < third{
			sum := nums[first] + nums[second] + nums[third]
			if sum == target {
				return target
			}
			update(sum)
			if sum - target > 0 {
				 third --
			} else {
				second++
			}
		}
	}

	return aim
}


func abs(a int) int {
	if a < 0 {
		return a * -1
	} else {return a}
}

func main() {
	arr := []int{1, 1, 2, 3}
	res := threeSumClosest(arr, 2)
	fmt.Println(res)
}
