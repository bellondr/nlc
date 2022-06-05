package main

import "sort"

type MyHeapInterface interface {
	sort.Interface
	Push(x interface{})
	Pop()interface{}
}

func Init(m MyHeapInterface) {
	l := m.Len()
	for i := l/2 - 1; i >= 0; i-- {
		down(m, i, l)
	}
}

func down(m MyHeapInterface, i, n int) {
	tmp := i
	for tmp * 2 < n {
		l := left(tmp)
		r := right(tmp)
		if r < n && m.Less(l, r) {
			l = r
		}
		if
	}
}
func parent(root int) int {
	return root/2
}
func left(root int) int {
	return root * 2
}
func right(root int) int {
	return root * 2 + 1
}

func less(x, y int) bool {
	return x < y
}
func swap(data []interface{}, x, y int)  {
	data[x], data[y] = data[y], data[x]
}
func swim(data []interface{}, x int) {
	if x > 1 && less(parent(x), x) {
		swap(data, parent(x), x)
		x = parent(x)
	}
}

func sink(data[]interface{}, x int) {
	for left(x) < len(data) {
		tmp := left(x)
		if right(x) < len(data) && less(tmp, right(x)) {
			tmp = right(x)
		}
		if less(tmp, x) {
			return
		}
		swap(data, x, tmp)
		x = tmp
	}
}