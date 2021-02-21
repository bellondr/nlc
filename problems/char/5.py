# package main
#
# import "fmt"
#
# func main() {
# 	fmt.Printf(longestPalindrome("cbbd"))
# }
# func longestPalindrome(s string) string {
# 	max := 0
# 	res := ""
# 	for i := 0; i < len(s); i++ {
# 		l := 1
# 		if l > max {
# 			max = l
# 			res = s[i:i+1]
# 		}
# 		extend := 1
# 		for {
# 			if i-extend >= 0 && i + extend < len(s) && s[i-extend]==s[i+extend] {
# 				l += 2
# 				extend++
# 			} else {
# 				break
# 			}
#
# 		}
# 		if l > max {
# 			max = l
# 			fmt.Println("extend: ", extend)
# 			fmt.Println("l: ", l)
# 			fmt.Println("index: ", i)
# 			res = s[i-extend+1:i+extend]
# 		}
# 		if i+1<len(s) && s[i]==s[i+1] {
# 			extend = 1
# 			l = 2
# 			for {
# 				if i+1+extend < len(s) && i-extend >= 0 && s[i-extend] == s[i+1+extend] {
# 					extend++
# 					l += 2
# 				} else {
# 					break
# 				}
# 			}
# 			if l > max {
# 				max = l
# 				res = s[i-extend+1:i+1+extend]
# 			}
# 		}
# 	}
#
# 	return res
# }
