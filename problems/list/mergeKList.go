package main

type ListNodeHeap []*ListNode

func (l ListNodeHeap)Len() int {
	return len(l)
}
func (l ListNodeHeap)Less(i, j int) bool {
	return l[i].Val < l[j].Val
}
func (l ListNodeHeap)Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l *ListNodeHeap)Push(x interface{}) {
	*l = append(*l, x.(*ListNode))
}
func (l *ListNodeHeap) Pop() interface{} {
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummyNode := &ListNode{}
	head := dummyNode
	lnh := &ListNodeHeap{}
	Init(lnh)
	for _, list := range lists {
		for list != nil {
			Push(lnh, list)
			list = list.Next
		}
	}


	for lnh.Len() > 0  {
		top := Pop(lnh).(*ListNode)
		head.Next = top
		head = head.Next
	}
	return dummyNode.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	n := len(lists)
	l := mergeKLists2(lists[:n/2])
	r := mergeKLists2(lists[n/2:])
	return mergeTwoLists(l, r)
}


type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i] == nil {
		return false
	}
	if pq[j] == nil {
		return true
	}
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}