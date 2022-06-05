package main

func reversePairs(nums []int) int {
	arr := make([]int, len(nums))
	tmp := make([]int, len(nums))
	for i := range nums {
		arr[i] = nums[i]
	}
	count := 0
	msort(arr, tmp, &count, 0, len(nums) - 1)
	return count
}

func msort(arr, tmp []int, count *int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right - left)/2
	msort(arr, tmp, count, left, mid)
	msort(arr, tmp, count, mid+1, right)
	mmerge(arr, tmp, count, left, mid, right)
}

func mmerge(arr, tmp []int, count *int, left, mid, right int) {
	for p := left; p <= right; p++ {
		tmp[p] = arr[p]
	}
	end := mid+1
	for i := left; i <= mid; i++ {
		for end <= right && int64(arr[i]) > int64(arr[end]) * 2 {
			end++
		}
		*count += (end - mid - 1)
	}

	l, r := left, mid+1
	for p:= left; p <= right; p++ {
		if l == mid+1 {
			arr[p] = tmp[r]
			r++
		} else if r == right + 1 {
			arr[p] = tmp[l]
			l++
		} else if tmp[l] > tmp[r] {
			arr[p] = tmp[r]
			r++
		} else {
			arr[p] = tmp[l]
			l++
		}
	}
}