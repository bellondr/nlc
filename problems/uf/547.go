package main

import (
	"fmt"
)

func main() {
	isConnected := [][]int{{1,1,0},{1,1,0},{0,0,1}}
	fmt.Println(findCircleNum(isConnected))
}
func findCircleNum(isConnected [][]int) int {
    m := len(isConnected)
    u := NewUF(m)
    for i := 0; i < m; i++ {
        for j:=0; j < m; j++ {
            if isConnected[i][j] == 1 {
                u.unin(i, j)
            }
        }
    }
    return u.cnt

}

func NewUF(n int) *UF {
    u := &UF{
        parent: make(map[int]int, n),
        cnt: n,
        size: make([]int, n),
    }
    for i := 0; i < n; i++ {
        u.parent[i] = i
    }
    for i := 0; i < n; i++ {
        u.size[i] = 1
    }
    return u
}

type UF struct {
    parent map[int]int
    cnt int
    size []int
}

func(u *UF) find(x int) int {
    for ;x != u.parent[x]; {
        x = u.parent[x]
    }
    return x
}

func (u *UF)connect(x, y int) bool {
    return u.find(x) == u.find(y)
}

func (u *UF)unin(x, y int) {
    if u.connect(x, y) {
        return
    }
    parentX := u.find(x)
    parentY := u.find(y)
    if u.size[parentX] > u.size[parentY] {
        u.parent[parentY] = parentX
        u.size[parentX] += u.size[parentY]
    } else {
        u.parent[parentX] = parentY
        u.size[parentY] += u.size[parentX]
    }
    u.cnt -= 1
}
