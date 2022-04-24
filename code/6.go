// Z字形变换
package main

import "fmt"

func convert(s string, numRows int) string{
	str := ""
	if numRows >= len(s) || numRows == 1{
		return s
	}
	temp := 0
	arr := make([] string, numRows)
	for i := numRows-1; i < len(s); i+=2*numRows-2{
		for num := 0; num < numRows; num++{
			arr[numRows-num-1] += string(s[i-num])
			if num > 0 && num < (numRows-1) && (i+num) < len(s) {
				arr[numRows-num-1] += string(s[i+num])
			}
		}
		if i + 2*numRows-2  >= len(s){
			temp = i+numRows-1
		}
	}
	for j := temp; j < len(s); j++ {
		arr[j-temp] += string(s[j])
	}
	for k := 0; k < numRows; k++ {
		str += arr[k]
	}
	return str
}

func main() {
	var s = "ABCDE"
	var numRows = 4
	res := convert(s, numRows)
	fmt.Println(res)
}
