// 二进制求和
package main

import (
	"fmt"
	"strconv"
)


func addBinary(a string, b string) string {
	var sum = ""
	flag := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 && j >= 0; i, j = i -1, j -1{
		// "0" = 48
		valueA, _ := strconv.Atoi(string(a[i]))
		valueB, _ := strconv.Atoi(string(b[j]))
		temp := valueA + valueB  + flag
		if temp >= 2 {
			sum = strconv.Itoa(temp-2) + sum
			flag = 1
			continue
		}else {
			sum = strconv.Itoa(temp) + sum
			flag = 0
		}
	}
	arr, space := longStr(a, b)
	for k := space -1; k >= 0; k--{
		valueK, _ := strconv.Atoi(string(arr[k]))
		if valueK + flag == 2 {
			flag = 1
			sum = "0"+ sum
		}else {
			sum = strconv.Itoa(valueK+flag) + sum
			flag = 0
		}
	}
	if flag == 1 {
		sum = "1" + sum
	}
	return sum
}

func longStr(a string, b string) (string, int){
	if len(a) > len(b){
		return a, len(a) - len(b)
	}else {
		return b, len(b) - len(a)
	}
}

func main() {
	var a, b = "1010", "1011"
	res := addBinary(a, b)
	fmt.Println(res)
}
