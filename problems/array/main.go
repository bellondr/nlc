package main

func main()  {
	testCheckIn()
}

func testSubarraySum() {
	rs := subarraySum([]int{1,2,3}, 3)
	print(rs)
}

func testMinWindow() {
	print(minWindow("ADOBECODEBANC", "ABC"))
}

func testCheckIn() {
	print(checkInclusion("trinitrophenylmethylnitramine", "razinetrinitrophenylmethylnitramine"))

	//print(checkInclusion("ab", "eidba"))
}