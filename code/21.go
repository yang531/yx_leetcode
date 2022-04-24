package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{Val:-1}
	prev := prehead
	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			prev.Next = l1
			l1 = l1.Next
		}else{
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil{
		prev.Next = l2
	}else{
		prev.Next = l1
	}
	return prehead.Next
}


func main() {
	var list1 ListNode
	var list2 ListNode
	mergeTwoLists(&list1, &list2)
}
