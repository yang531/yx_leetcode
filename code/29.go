// 两数相除
package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	flag := false
	res := 0
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			 return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend > 0 {
		dividend = -dividend
		flag = !flag
	}
	if divisor > 0 {
		divisor = -divisor
		flag = !flag
	}

	left, right := 1, -dividend
	for left <= right {
		mid := left + (right - left)>>1
		if add(mid, divisor) > dividend && add(mid + 1, divisor) >= dividend{
			left = mid + 1
		} else if add(mid, divisor) < dividend {
			right = mid - 1
		} else {
			res = mid
			break
		}
	}
	if flag {
		return -res
	}else {
		return res
	}
}


func add(x int, y int) int{
	sum := 0
	for i := 0; i < x; i++ {
		sum += y
	}
	return sum
}

func main() {
	res := divide(10000000000,1)
	fmt.Println(res)
}
