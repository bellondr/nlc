package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	rs := Constructor()
	rs.Insert(0)
	rs.Insert(1)
	rs.Remove(0)
	rs.Insert(2)
	rs.Remove(1)

	fmt.Println(rs.GetRandom())
}


type RandomizedSet struct {
	dataMap map[int]int
	dataArr []int
}


func Constructor() RandomizedSet {
	return RandomizedSet{
		dataMap: make(map[int]int),
		dataArr: make([]int, 0),
	}
}


func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.dataMap[val]
	if ok {
		return !ok
	}
	this.dataMap[val] = len(this.dataArr)
	this.dataArr = append(this.dataArr, val)
	return !ok
}


func (this *RandomizedSet) Remove(val int) bool {
	fmt.Printf("val is: %d \n", val)
	index, ok := this.dataMap[val]
	if ok {
		fmt.Printf("arr: %v \n", this.dataArr)
		fmt.Printf("index: %d \n", index)
		tail := len(this.dataArr) - 1
		fmt.Printf("tail: %d \n", tail)
		this.dataMap[this.dataArr[tail]] = index
		delete(this.dataMap, val)
		this.dataArr[index], this.dataArr[tail] = this.dataArr[tail], this.dataArr[index]
		this.dataArr = this.dataArr[:tail]
		fmt.Printf("arr: %v \n", this.dataArr)
	}
	return ok
}


func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.dataArr))
	return this.dataArr[index]
}
