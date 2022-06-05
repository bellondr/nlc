package main

type Interface interface {
	Push(x any)
	Pop() any
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Init(h Interface) {
	n := h.Len()
	for i := n/2 - 1; i >=0; i-- {
		down(h, i, n)
	}
}

func Push(h Interface, x any) {
	h.Push(x)
	up(h, h.Len() - 1)
}
func Pop(h Interface) any {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func down(h Interface, index, n int) bool {
	i := index
	for {
		j1 := 2 * i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > index
}

func up(h Interface, j int) {
	for {
		i := (j - 1) / 2
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}