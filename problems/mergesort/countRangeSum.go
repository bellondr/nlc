package main


func countRangeSum(nums []int, lower int, upper int) int {
	preSum := make([]int64, len(nums)+1)
	tmp := make([]int64, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + int64(nums[i])
	}
	count := 0
	sort(preSum, tmp, &count, 0, len(preSum) - 1, int64(lower), int64(upper))
	return count
}

func sort(preSum, tmp []int64, count *int, le, ri int, l, u int64) {
	if le >= ri {
		return
	}
	mid := le + (ri - le)/2
	sort(preSum, tmp, count, le, mid, l, u)
	sort(preSum, tmp, count, mid+1, ri, l, u)
	cmerge(preSum, tmp, count, le, mid, ri,l, u)
}

func cmerge(preSum, tmp []int64, count *int, le, mid, ri int, l, u int64) {
	for p := le; p <= ri; p++{
		tmp[p] = preSum[p]
	}
	start, end := mid + 1, mid + 1
	for i := le; i <= mid; i++ {
		for start <= ri && preSum[start] - preSum[i] < l {
			start++
		}
		for end <= ri && preSum[end] - preSum[i] <= u {
			end++
		}
		*count += (end - start)
	}

	ls, rs := le, mid+1
	for p := le; p <= ri; p++ {
		if ls == mid+1 {
			preSum[p] = tmp[rs]
			rs++
		} else if rs == ri + 1 {
			preSum[p] = tmp[ls]
			ls++
		} else if tmp[rs] > tmp[ls] {
			preSum[p] = tmp[ls]
			ls++
		} else {
			preSum[p] = tmp[rs]
			rs++
		}
	}
}

