package main
import "fmt"

func testLongestPalindrome() {
	s := "ac"
	fmt.Println(longestPalindrome(s))
}


func longestPalindrome(s string) string {
	maxSt := ""
	for i := range s {
		t := Palindrome(s, i, i)
		if len(t) > len(maxSt) {
			maxSt = t
		}
		t = Palindrome(s, i, i+1)
		if len(t) > len(maxSt) {
			maxSt = t
		}
	}
	return maxSt
}

func Palindrome(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1:j]
}

