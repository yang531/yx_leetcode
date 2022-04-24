// 字符串转换整数
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	n := len(s)
	if n == 0 {
		return 0
	}
	// 0-9 ASII 48-57
	if string(s[0]) != "-" && string(s[0]) != "+" && (s[0] > 57 || s[0] < 48)  {
		return 0
	}
	for i := 0; i < n; i++{
		if i == 0 && (string(s[i]) == "-" || string(s[i]) == "+"){
			continue
		}
		if s[i] >= 48 && s[i] <= 57{
			continue
		}
		s = s[0:i]
		break
	}
	res, _ := strconv.Atoi(s)
	if res < math.MinInt32 {
		res = math.MinInt32
	} else if res > math.MaxInt32{
		res = math.MaxInt32
	}
	return res
}

func main() {
	s := "-5-"
	res := myAtoi(s)
	fmt.Println(res)
}
