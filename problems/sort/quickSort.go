package main

import (
	"fmt"
)

func main() {
	arr := []int{4,5,3,2,7}
	quickSort(arr,0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int){
	if left < right {
		i, j := left, right
		key := arr[(left+right)/2]
		for i < j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if  i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if left < j {
			quickSort(arr, left, j)
		}
		if right > i {
			quickSort(arr, i, right)
		}
	}
}
