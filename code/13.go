package main

import (
	"fmt"
)


func romanToInt(s string) int {
	res := 0
	for i:=len(s)-1; i >=0; i--{
		val := string(s[i])
		switch val {
		case "I":
			if i == len(s) -1{
				res += 1
			} else {
				temp := string(s[i+1])
				if temp != "V" && temp != "X"{
					res += 1
				} else {
					res -= 1
				}
			}
		case "V": res += 5
		case "X":
			if i == len(s) - 1 {
				res += 10
			} else {
				temp := string(s[i+1])
				if temp != "L" && temp != "C"{
					res += 10
				} else {
					res -= 10
				}
			}
		case "L": res += 50
		case "C":
			if i == len(s) -1 {
				res += 100
			} else {
				temp := string(s[i+1])
				if temp != "D" && temp != "M"{
					res += 100
				} else {
					res -= 100
				}
			}
		case "D": res += 500
		case "M": res += 1000
		}
	}
	return res
}

func main()  {
	res := romanToInt("MCMXCIV")

	fmt.Println(res)
}

