package main

import "fmt"

func mySqrt(x int) int {
	l, r := 0, x
	res := 0
	for l <= r {
		mid := (l+r) / 2
		if mid * mid <= x {
			l = mid + 1
			res = mid
		} else {
			r = mid - 1
		}
	}
	return res
}


func main() {
	res := mySqrt(10)
	fmt.Println(res)
}
