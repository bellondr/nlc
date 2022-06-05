package main

import "container/list"

type LFUCache struct {
	nodes map[int]*list.Element
	lists map[int]*list.List
	capacity int
	min int
}

type node struct {
	key int
	value int
	freq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{nodes: make(map[int]*list.Element),
		lists: make(map[int]*list.List),
		capacity: capacity,
		min: 0,
	}
}


func (this *LFUCache) Get(key int) int {
	v, ok := this.nodes[key]
	if !ok {
		return -1
	}
	n := v.Value.(*node)
	this.lists[n.freq].Remove(v)
	n.freq++
	if _, ok := this.lists[n.freq]; !ok {
		this.lists[n.freq] = list.New()
	}
	newlist := this.lists[n.freq]
	newnode := newlist.PushBack(n)
	this.nodes[key] = newnode
	if n.freq - 1 == this.min && this.lists[n.freq - 1].Len() == 0 {
		this.min++
	}
	return n.value
}


func (this *LFUCache) Put(key int, value int)  {
	if this.capacity == 0 {
		return
	}

	if v, ok := this.nodes[key]; ok {
		n := v.Value.(*node)
		n.value = value
		this.Get(key)
		return
	}

	if this.capacity == len(this.nodes) {
		list := this.lists[this.min]
		frontNode := list.Front()
		delete(this.nodes, frontNode.Value.(*node).key)
		list.Remove(frontNode)
	}

	this.min = 1
	n := &node{key: key,
		value: value,
		freq: 1,
	}

	if _, ok := this.lists[1]; !ok {
		this.lists[1] = list.New()
	}

	list1 := this.lists[1]
	newnode := list1.PushBack(n)
	this.nodes[key] = newnode
}
