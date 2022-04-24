// 无重复字符的最长子串
package main

import (
	"fmt"
)

// 滑窗法
func lengthOfLongestSubstring(s string) int {
	var length = len(s)
	if length <= 1 {
		return len(s)
	}
	var left, right, maxNum = 0, 0, 1
	var windows  = make(map[byte]bool)
	for right < length{
		for windows[s[right]] {
			delete(windows, s[left])
			left++
		}
		windows[s[right]] = true
		if right - left + 1 > maxNum {
			maxNum = right - left + 1
		}
		right++
	}
	return maxNum
}


func main() {
	var str = ""
	res := lengthOfLongestSubstring(str)
	fmt.Println(res)
}
