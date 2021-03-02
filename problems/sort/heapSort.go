package main

import "fmt"

func main() {
	arr := []int{1,4,7,2,4,3}
	heapSort(arr)
	fmt.Println(arr)
}

func heapSort(arr[]int) {
	N := len(arr) - 1

	for k := N/2; k >= 0; k-- {
		sink(arr, k, N)
	}

	for N > 0 {
		arr[N], arr[0] = arr[0], arr[N]
		N--
		sink(arr, 0, N)
	}

}

func sink(arr []int, k, N int) {
	for {
		i := k * 2
		if i > N {
			break
		}
		if i < N && arr[i+1] > arr[i] {
			i = i+1
		}
		if arr[k] >=  arr[i] {
			break
		}
		arr[i], arr[k] = arr[k], arr[i]
		k = i
	}
}
