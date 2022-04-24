// 电话号码的字母组合——回溯
package main

import (
	"fmt"
)

var phone = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var arr []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	arr = [] string{}
	recursion(digits, 0, "")
	return arr
}

// 递归
func recursion(digits string, num int, content string){
	if num == len(digits) {
		arr = append(arr, content)
	} else {
		temp := phone[string(digits[num])]
		for i := 0; i < len(temp); i++ {
			// 递归
			recursion(digits, num+1, content + string(temp[i]))
		}
	}
}

func main() {
	res := letterCombinations("1234")
	fmt.Println(res)
}
