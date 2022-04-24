// 罗马数字转整数
package main

import (
	"fmt"
)


func romanToInt1(s string) int {
	mMap := map[string] int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	for i:=len(s)-1; i >=0; i--{
		val := string(s[i])
		if i == len(s) - 1 {
			res += mMap[val]
		} else {
			temp := string(s[i+1])
			if val == "I" {
				if temp != "V" && temp != "X"{
					res += mMap[val]
				} else {
					res -=  mMap[val]
				}
			} else if val == "X" {
				if temp != "L" && temp != "C"{
					res += mMap[val]
				} else {
					res -= mMap[val]
				}
			} else if val == "C"{
				if temp != "D" && temp != "M"{
					res += mMap[val]
				} else {
					res -= mMap[val]
				}
			} else {
				res += mMap[val]
			}
		}

	}
	return res
}

func main()  {
	res := romanToInt1("MCMXCIV")

	fmt.Println(res)
}

