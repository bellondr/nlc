package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	return reverseK(head, k)
}

func reverseK(head *ListNode, k int) *ListNode {
	var pre *ListNode
	cur := head
	i := 0
	for tmp := head; tmp != nil && i < k; i++ {
		tmp = tmp.Next
	}
	if i == k {
		for i = 0; i < k; i++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}

	} else {
		return head
	}

	if cur != nil {
		head.Next = reverseK(cur, k)
	}
	return pre
}
