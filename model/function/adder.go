package main

import (
	"fmt"
)

func Adder()func(int) int {
	num := 0
	return func(i int) int {
		fmt.Println(num)
		num += i
		return num
	}
}

func testAdder() {

	adder := Adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %d, num: %d \n", i, adder(i))
	}
}

func fibo()func()int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func TestFibo()  {
	fi := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(fi())
	}
}

func main()  {
	TestFibo()
}
