package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s)-1; i++{
		left1, right1 := expand(s, i, i)
		left2, right2 := expand(s, i, i+1)
		if right1 - left1 > end - start{
			end = right1
			start = left1
		}
		if right2 - left2 > end - start{
			end = right2
			start = left2
		}
	}
	return s[start:end+1]
}

func expand(s string, left int, right int) (int, int)  {
	for ; left>=0 && right<len(s) && s[left]==s[right] && left<=right; left,right=left-1, right+1{}
	return left+1, right-1
}

func main() {
	s := "a"
	res := longestPalindrome(s)
	fmt.Println(res)
	
}
