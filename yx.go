package main

import (
	"fmt"
	"strconv"
)

func main()  {
	arr := "-012344"
	fmt.Println(arr[0])
	fmt.Println(int(arr[0]))
	fmt.Println(string(arr[0]))
	fmt.Println(strconv.Atoi(arr))
	fmt.Println("-----------------")

	arr1 := []string{"1", "2", "3"}
	fmt.Println(arr1[0])
	fmt.Println(strconv.Atoi(arr1[0]))

	fmt.Println(string(1))
	fmt.Println(strconv.Itoa(1))
}

