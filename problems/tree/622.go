package main

type MyCircularQueue struct {
	data []int
	cap int
}


func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data: []int{},
		cap: k,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.data) < this.cap {
		this.data = append(this.data, value)
		return true
	}
	return false
}


func (this *MyCircularQueue) DeQueue() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[1:]
	return true
}


func (this *MyCircularQueue) Front() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[0]
}


func (this *MyCircularQueue) Rear() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.data) == 0
}


func (this *MyCircularQueue) IsFull() bool {
	return len(this.data) == this.cap
}
