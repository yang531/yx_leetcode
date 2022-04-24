package main

import "fmt"

// 这种方法受int类型的上限所致，无法计算过长的数组
//func plusOne1(digits []int) []int {
//	sum := 0
//	var res []int
//	for _, value := range digits{
//		sum = sum*10 + value
//	}
//	sum += 1
//	for {
//		temp := sum % 10
//		res = append([]int{temp}, res...)
//		sum /= 10
//		if sum == 0 {
//			break
//		}
//	}
//	return res
//}

func plusOne(digits []int) []int {
	for i := len(digits) -1; i >=0; i-- {
		digits[i] += 1
		if digits[i] != 10 {
			break
		}else {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
				return digits
			}
			continue
		}
	}
	return digits
}


func main() {
	digits := [] int{9, 9, 9, 9, 9}
	res := plusOne(digits)
	fmt.Println(res)
}
