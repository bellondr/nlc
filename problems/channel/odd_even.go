package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	//oddPrint()
	//myOE()
	mytest()
}

func odPC()  {
	ch := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("Send:", i)
			time.Sleep(time.Second)
			ch <- i
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			v := <-ch
			fmt.Println("Receive:", v)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(20 * time.Second)
}

func mytest()  {
	ch := make(chan int)
	//go func() {
	//	select {
	//	case <- ch:
	//		fmt.Println("pop channel done")
	//	}
	//}()
	go func() {
		ch <- 1
		fmt.Println("add channel done")
	}()
	time.Sleep(10 * time.Second)
}
func myOE() {
	ch := make(chan int)

	go func() {
		index := 1
		for i := 0; i < 20; i++ {
			fmt.Println("send: ", index)
			ch <- i
			index += 2
		}
	}()

	go func() {
		index := 2
		for i := 0; i < 20; i++ {
			<- ch
			fmt.Println("receive:", index)
			time.Sleep(time.Second)
			index += 2
		}
	}()
	time.Sleep(10 * time.Second)
}

func oddPrint() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 0
	var wg sync.WaitGroup
	wg.Add(2)
	even := func(ch1, ch2 chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i <= 100; i += 2 {
			select {
			case <-ch1:
				fmt.Println(i)
			}
			ch2 <- 0
		}
	}
	odd := func(ch1, ch2 chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			select {
			case <- ch2:
				fmt.Println(i)
			}
			ch1 <- 0
		}
	}

	go even(ch1, ch2, &wg)
	go odd(ch1, ch2, &wg)

	wg.Wait()
}

func printEven(ch <-chan int) {
	for i := 1; i <= 100; i ++ {
		v := <- ch
		fmt.Println("recive: ", v )
	}
}

func printOdd(ch chan<- int) {
	for i := 1; i <= 100; i ++ {
		ch <- i
		fmt.Println("send: ", i)
	}
}
func printWithOneChannel() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 100; i +=2 {
			v := <- ch
			fmt.Println(v)
		}
	}()
	go func() {
		for i := 1; i <= 100; i += 2 {
			if i % 2 == 1{
				ch <- i
			} else {
				fmt.Println(i)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
