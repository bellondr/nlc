package main


type NumArray struct {
	sumArray []int
}


func Constructor(nums []int) NumArray {
	sumArray := make([]int, len(nums)+1)
	for i, v := range nums {
		sumArray[i+1] = sumArray[i]+v
	}
	return NumArray { sumArray: sumArray     }
}


func (this *NumArray) SumRange(left int, right int) int {
	return this.sumArray[right + 1] - this.sumArray[left]
}