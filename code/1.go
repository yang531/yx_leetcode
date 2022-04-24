package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int)
	for i := range nums {
		if v, ok := m[target-nums[i]]; ok && v!= i {
			res = append(res, v, i)
			break
		}
		m[nums[i]] = i
	}
	return res
}

func main()  {
	nums := []int{3, 3}
	res := twoSum(nums, 6)
	fmt.Println(res)
}

