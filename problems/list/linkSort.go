package main

type ListNode struct {
	Val int
	Next *ListNode
}


func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	l := sortList(head)
	r := sortList(slow)
	return mergeList(l, r)
}

func mergeList(left, right *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	for left != nil && right != nil {
		if left.Val > right.Val {
			cur.Next = right
			right = right.Next
		} else {
			cur.Next = left
			left = left.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return root.Next
}

