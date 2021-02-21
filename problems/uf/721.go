package main

import (
	"fmt"
)

func main() {
	accounts := [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}
	fmt.Println(accountsMerge(accounts))
}

func accountsMerge(accounts [][]string) [][]string {
	uf := NewAccountUF()
	emailName := make(map[string]string)
	for _, account := range accounts {
		for i := range account {
			if i == 0 {
				continue
			}
			emailName[account[i]] = account[0]
			if i < len(account) - 1 {
				uf.Unin(account[i], account[i+1])
			}
		}
	}
	resMap := make(map[string][]string)
	for k := range emailName {
		root := uf.find(k)
		_, ok := resMap[root]
		if ok {
			resMap[root] = append(resMap[root], k)
		} else {
			resMap[root] = []string{k}
		}
	}
	res := make([][]string, 0)
	for k , v := range resMap {
		sortStringArray(v)
		res = append(res, append([]string{emailName[k]}, v...))
	}
	return res
}

func sortStringArray(arr []string){
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
type AccountUF struct {
	partent map[string]string
}

func NewAccountUF() *AccountUF {
	return &AccountUF{
		partent: make(map[string]string),
	}
}

func (a *AccountUF) find(name string) string {
	for {
		v, ok := a.partent[name]
		if !ok {
			a.partent[name] = name
			break
		}
		if v == name {
			break
		}
		name = v
	}
	return name
}

func (a *AccountUF) Unin(n1, n2 string) {
	a1 := a.find(n1)
	a2 := a.find(n2)
	a.partent[a1] = a2
}
//func accountsMerge(accounts [][]string) [][]string {
//	m := len(accounts)
//	releation := make([][]int, m)
//	for i := range releation {
//		releation[i] = make([]int, m)
//	}
//	for i := 0; i < m; i++ {
//		releation[i][i] = 1
//		for j := i+1; j < m; j++ {
//			if neighbor(accounts[i][1:], accounts[j][1:]) {
//				releation[i][j] = 1
//			}
//		}
//	}
//	fmt.Println(releation)
//	res := make([][]string, 0)
//	tag := make([]bool, m)
//	for i := 0; i < m; i++ {
//		if tag[i] {
//			continue
//		}
//		tract := make(map[string]bool)
//		for j := 0; j < m; j ++ {
//			if releation[i][j] == 1 {
//
//			}
//			releation[i][j] = 0
//		}
//
//	}
//	return res
//}
//
//func convertRes(m map[string]bool, name string) []string {
//	res := make([]string, 0)
//	res = append(res, name)
//	for k, v := range m {
//		if v {
//			res = append(res, k)
//		}
//	}
//	return res
//}
//
//func collectEmails(res map[string]bool, accounts [][]string, rel [][]int, x, y int)  {
//	if x >= len(rel) || x < 0 || y >= len(rel) || y < 0 {
//		return
//	}
//	for i := range accounts[x] {
//		if i > 0 {
//			res[accounts[x][i]] = true
//		}
//	}
//	for i := range accounts[y] {
//		if i > 0 {
//			res[accounts[y][i]] = true
//		}
//	}
//	rel[x][y] = 0
//	if x > 0 && rel[x-1][y] == 1 {
//		collectEmails(res, accounts, rel, x-1, y)
//	}
//	if x < len(rel) - 1 && rel[x+1][y] == 1 {
//		collectEmails(res, accounts, rel, x+1, y)
//	}
//	if y < len(rel) - 1 && rel[x][y+1] == 1 {
//		collectEmails(res, accounts, rel, x, y+1)
//	}
//	if y > 0 && rel[x][y-1] == 1 {
//		collectEmails(res, accounts, rel, x, y-1)
//	}
//}
//
//
//func neighbor(p1[]string, p2 []string) bool {
//	for _, p := range p1 {
//		for _, n := range p2 {
//			if p == n {
//				return true
//			}
//		}
//	}
//	return false
//}
