package main

import (
	"fmt"
)

func main() {
	case1 := []string{"c==c","b==d","x!=z"}
	fmt.Println(equationsPossible(case1))
}

func equationsPossible(equations []string) bool {
	uf := NewEquationUF()
	for _, eq := range equations {
		if eq[1:3] == "==" {
			uf.union(int(eq[0]), int(eq[3]))
		}
	}
	for _, eq := range equations {
		if eq[1:3] == "!=" {
			if uf.find(int(eq[0])) == uf.find(int(eq[3])) {
				return false
			}
		}
	}
	return true

}

func NewEquationUF() *EquationUF {
	return &EquationUF{
		parent: make(map[int]int),
	}
}

type EquationUF struct {
	parent map[int]int
}

func (e *EquationUF) find(s int) int {
	for {
		v, ok := e.parent[s]
		if !ok {
			e.parent[s] = s
			break
		}
		if v == s {
			break
		}
		s = v
	}
	return s
}

func (e *EquationUF) connect(s1, s2 int) bool {
	return e.find(s1) == e.find(s2)
}

func (e *EquationUF) union(s1, s2 int) {
	if e.connect(s1, s2) {
		return
	}
	e.parent[e.find(s1)] = e.find(s2)
}
