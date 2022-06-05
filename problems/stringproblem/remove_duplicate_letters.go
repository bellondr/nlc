package main

import "fmt"

func testRemoveDuplicateLetters() {
	fmt.Println(removeDuplicateLetters("bcabc"))
}

func removeDuplicateLetters(s string) string {
	// count, instack, array
	sArr := []byte(s)
	count := make([]byte, 256)
	for _, v := range s {
		count[v]++
	}
	instack := make([]bool, 256)
	array := make([]byte, 0)
	for _, v := range sArr {
		count[v]--
		if instack[v] {
			continue
		}
		for len(array) > 0 {
			tail := len(array) - 1
			if array[tail] <= v {
				break
			}
			val := array[tail]
			if count[val] == 0 {
				break
			}
			instack[val] = false
			array = array[:tail]
		}
		array = append(array, v)
		instack[v] = true
	}
	return string(array)
}