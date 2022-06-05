package main

func main() {
	testSplitArray()
}

func testShip()  {
	//wi := []int{5,5,5,5,5,5,5,5,5,5}

	wi2 := []int{1,2,3,4,5,6,7,8,9,10}
	print(shipWithinDays(wi2, 5))
}

func shipWithinDays(weights []int, days int) int {
	m := weights[0]
	for i := range weights {
		if m < weights[i] {
			m = weights[i]
		}
	}
	left, right := 1, m * len(weights)
	for left < right {
		mid := left + (right - left)/2
		d := converyorDays(weights, mid, days)
		if d > 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func converyorDays(weights []int, n int, days int) int {
	day := 1
	cap := n
	for i := range weights {
		if cap >= weights[i] {
			cap -= weights[i]
		} else {
			day++
			cap = n - weights[i]
			if cap < 0 {
				return 1
			}
		}
	}
	return day - days
}
