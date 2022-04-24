// 有效的括号
package main

import (
	"container/list"
	"fmt"
)


func isValid(s string)  bool{
	//生成一个双向链表当栈
	stack := list.New()
	for _, value := range s{
		if string(value) != ")" && string(value) != "}" && string(value) != "]"{
			//入栈
			stack.PushBack(value)
		} else if stack.Back() == nil  {
			return false
		} else{
			temp := stack.Back().Value.(int32)
			if value - temp > 2 || value - temp < 0{
				return false
			} else {
				//出栈
				stack.Remove(stack.Back())
			}
		}
	}
	if stack.Back() != nil{
		return false
	}
	return true
}


func main()  {
	//strs := [] string{"dog", "racecar", "car"}
	res :=isValid("(){}[]")

	fmt.Println(res)
}

