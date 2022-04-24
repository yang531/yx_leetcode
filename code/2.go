// 两数相加 ——待定
package main

import "fmt"

func makeSlice()  {
	mSlice := make([] string, 3)
	mSlice[0] = "yang"
	mSlice[1] = "liang"
	mSlice[2] = "hai"
	fmt.Println(mSlice)
}

func makeMap()  {
	mMap := make(map[string] string)
	mMap["name"] = "杨"
	mMap["age"] = "123"
	fmt.Println(mMap)
}

func main()  {
	makeSlice()
	makeMap()
}