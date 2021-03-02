package main

import "fmt"

func main() {
	a := 20
	it := Item{10, &a, Detail{
		Val: 30,
	}, &Detail{
		Val: 40,
	}}
	fmt.Printf("%d, %d, %d, %d \n", it.v, *it.pv, it.de.Val, it.pde.Val)
	it.Do()
	fmt.Printf("%d, %d, %d, %d \n", it.v, *it.pv, it.de.Val, it.pde.Val)
	it.PreDo()
	fmt.Printf("%d, %d, %d, %d \n", it.v, *it.pv, it.de.Val, it.pde.Val)
}

type Detail struct {
	Val int
}

type Item struct {
	v int
	pv *int
	de Detail
	pde *Detail
}

func(i Item)Do() {
	fmt.Println(i)
	i.v = 11
	*(i.pv) = 21
	i.de.Val = 31
	(*i.pde).Val = 41
}

func (i *Item)PreDo() {
	fmt.Println(i)
	i.v = 19
	*(i.pv) = 29
	i.de.Val = 39
	(*i.pde).Val = 49
}
