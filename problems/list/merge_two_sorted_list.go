package main

import "fmt"

func main() {
	list1 :=&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	//list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	//list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8}}}
	l := reverseKGroup(list1, 2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := &ListNode{}
	head := root
	for ;list1 != nil && list2 != nil; {
		if list1.Val > list2.Val {
			head.Next = list2
			list2 = list2.Next
		} else {
			head.Next = list1
			list1 = list1.Next
		}
		head = head.Next
	}
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	return root.Next
}
