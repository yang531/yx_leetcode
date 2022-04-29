// 搜索二维矩阵
/*
先用二分法寻找可能出现在第几行，然后再判断这行中是否存在target
*/
package main

import (
	"fmt"
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	left, right, ans := 0, m-1, 0
	mid := 0
	for left <= right{
		mid = (left + right) >> 1
		if target < matrix[mid][0] {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	index := sort.SearchInts(matrix[ans], target)
	return index < len(matrix[0]) && matrix[ans][index] == target
}


func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	fmt.Println(searchMatrix(matrix, 8))
}
