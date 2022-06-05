package main

import "container/list"

type Deque struct {
	data *list.List
}

func NewDeque() *Deque {
	return &Deque{data:list.New()}
}

func (q *Deque) Push(val int) {
	for q.data.Len() != 0 && val > q.data.Back().Value.(int) {
		q.data.Remove(q.data.Back())
	}
	q.data.PushBack(val)
}

func (q *Deque) Pop(val int) {
	if q.data.Len() > 0 && q.Max() == val {
		q.data.Remove(q.data.Front())
	}
}

func (q *Deque) Max() int {
	return q.data.Front().Value.(int)
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	window := NewDeque()
	for i, v := range nums {
		if i < k - 1 {
			window.Push(v)
		} else {
			window.Push(v)
			res = append(res, window.Max())
			window.Pop(nums[i-k+1])
		}
	}
	return res
}
