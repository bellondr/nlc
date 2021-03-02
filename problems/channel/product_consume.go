package main

import (
	"fmt"
	"time"
	"sync"
)
var mu sync.Mutex
var res int

func main()  {
	ufPC()
}

func testCP() {
	ch := make(chan int)
	go produce1(ch)
	go consumer1(ch)
	time.Sleep(30*time.Second)
}
func myCP()  {
	ch := make(chan int, 3)
	go consumer(ch)
	go produce(ch)
	time.Sleep(5 * time.Second)
}

func consumer(ch chan int)  {
	for {
		v := <- ch
		fmt.Println("consume: ", v)
		time.Sleep(time.Second)
	}
}

func produce(ch chan int)  {
	index := 1
	for {
		ch <- index
		fmt.Println("produce: ", index)
		index++
	}
}

func ufPC()  {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("Send:", i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v := <-ch
			fmt.Println("Receive:", v)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(20 * time.Second)
}
func produce1(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Send:", i)
	}
}

func consumer1(ch <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println("Receive:", v)
		time.Sleep(time.Second)
	}
}
