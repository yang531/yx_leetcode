// 删除有序数组中的重复项
package main

import (
	"fmt"
)


func  removeDuplicates(nums []int) int {
	i := 0
	var j int
	for {
		if i >= len(nums) {
			break
		}

		for j = i+1; j < len(nums) && nums[i] == nums[j]; j++{
		}
		nums = append(nums[: i+1], nums[j:]...)
		i++
	}
	fmt.Println(nums)
	return len(nums)
}


func main()  {
	nums := [] int{0, 1, 1, 2}
	res := removeDuplicates(nums)
	fmt.Println(res)
}

