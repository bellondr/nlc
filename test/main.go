package main

import "fmt"

func main() {
	aar := []int{1,2,3,4,5,6,7,8}
	b := aar[3:]
	b[0] = 100
	fmt.Println(aar)
	fmt.Println(b)
}
