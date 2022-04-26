// 下一个排列
/*
（5,2,4,3,1）
1、从右往左找，先找出不符合降序的数字2， 再从右往左寻找比2大的数字3，
2、交换2， 3 -》(5,3,4,2,1)
3、将3后面的所有数字按照升序排列 -》（5,3,1,2,4）
->(5,)
*/
package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) []int {
	length := len(nums)
	min, max := 0, 0
	for i := length - 1; i >= 0; i--{
		if i == 0 {
			sort.Ints(nums)
			return nums
		}
		if nums[i-1] >= nums[i]{
			continue
		} else {
			min = i - 1
			for j := length -1; j >= 0; j--{
				if nums[j] <= nums[min] {
					continue
				} else {
					max = j
					break
				}
			}
			break
		}
	}
	nums[min], nums[max] = nums[max], nums[min]

	tempArr := nums[min+1:]
	sort.Ints(tempArr)
	res := append(nums[:min+1], tempArr...)
	return res
}



func main() {
	arr := []int{5,2,4,3,1}
	res := nextPermutation(arr)
	fmt.Println(res)
}
