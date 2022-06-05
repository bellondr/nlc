package main

type node struct {
	pre *node
	next *node
	key int
	val int
}

type LRUCache struct {
	hashMap map[int]*node
	head *node
	end *node
	capacity int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		hashMap: make(map[int]*node),
		capacity: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	v, ok := this.hashMap[key]
	if ok {
		if v != this.head {
			pre := v.pre
			next := v.next
			if pre != nil {
				pre.next = next
			}
			if next != nil {
				next.pre = pre
			}

			v.pre = nil
			v.next = this.head
			if v == this.end {
				this.end = pre
			}
			this.head.pre = v
			this.head = v
		}
		return v.val
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	v, ok := this.hashMap[key]
	if ok {
		v.val = value
		if v != this.head {
			pre := v.pre
			next := v.next
			if pre != nil {
				pre.next = next
			}
			if next != nil {
				next.pre = pre
			}
			if v == this.end {
				this.end = pre
			}
			v.pre = nil
			v.next = this.head
			this.head.pre = v
			this.head = v
		}
	} else {
		v = &node{
			key: key,
			val: value,
		}
		v.next = this.head
		if this.head != nil {
			this.head.pre = v
		} else {
			this.end = v
		}
		this.head = v
		this.hashMap[key] = v
	}
	if len(this.hashMap) >this.capacity {
		dk := this.end.key
		delete(this.hashMap, dk)
		pre := this.end.pre
		this.end.pre = nil
		pre.next = nil
		this.end = pre
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
