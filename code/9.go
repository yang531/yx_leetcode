package main

import (
	"fmt"
)


func isPalindrome(x int) bool {
	res:=0
	temp := x
	if x < 0{
		return false
	}
	for x!=0 {
		res = res*10 + x%10
		x /= 10
	}
	return res == temp
}

func main()  {
	res := isPalindrome(12321)

	fmt.Println(res)
}

