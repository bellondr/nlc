package main

type Pair struct {
	Val int
	Index int
}

func countSmaller(nums []int) []int {
	pairs := make([]Pair, len(nums))
	for i := range nums {
		pairs[i] = Pair{
			Val: nums[i],
			Index: i,
		}
	}

	count := make([]int, len(nums))

	mergeSort(pairs, count, 0, len(nums)-1)
	return count
}

func mergeSort(pairs []Pair, count []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	mergeSort(pairs, count, l, mid)
	mergeSort(pairs, count, mid+1, r)
	merge(pairs, count, l, mid, r)
}

func merge(pairs []Pair, count []int, l, mid, r int) {
	tmp := make([]Pair, r-l+1)
	for i := l; i <= r; i++{
		tmp[i] = pairs[i]
	}
	i, j := l, mid+1
	for p := l; p <= r; p++ {
		if i == mid+1 {
			pairs[p] = tmp[j]
			j++
		} else if j == r + 1 {
			pairs[p] = tmp[i]
			i++
			count[pairs[p].Index] += j - mid - 1
		} else if tmp[i].Val > tmp[j].Val {
			pairs[p] = tmp[j]
			j++
		} else {
			pairs[p] = tmp[i]
			i++
			count[pairs[p].Index] += j - mid - 1
		}
	}
}
