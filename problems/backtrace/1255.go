package main

import "fmt"

func main() {
	words := []string{"dog","cat","dad","good"}
	letters := []byte{'a','a','c','d','d','d','g','o','o'}
	score := []int{1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0}

	fmt.Println(maxScoreWords(words, letters, score))
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	lettersMap := make(map[int]int)
	for _, val := range letters {
		v, ok := lettersMap[int(val - 'a')]
		if !ok {
			lettersMap[int(val - 'a')] = 1
		} else {
			lettersMap[int(val - 'a')] = v + 1
		}
	}
	res := 0
	maxBacktrace(0, 0, &res, words, lettersMap, score)
	return res
}


func maxBacktrace(start, cur int, res *int, words []string, letters map[int]int, scores []int) {
	if start > len(words) {
		return
	}
	if cur > *res {
		*res = cur
	}
	for i := start; i < len(words); i++ {
		tmpMap := copyMap(letters)
		if removeWordInLetters(words[i], tmpMap) {
			maxBacktrace(i+1, cur + getWordScore(words[i], scores), res, words, tmpMap, scores)
		}
	}
}

func copyMap(letter map[int]int) map[int]int {
	res := make(map[int]int)
	for k, v := range letter {
		res[k] = v
	}
	return res
}

func getWordScore(word string, scores[]int) int{
	res := 0
	for i := range word {
		res += scores[word[i] - 'a']
	}
	return res
}

func removeWordInLetters(word string, letters map[int]int) bool {
	wm := make(map[int]int)
	for i := range word {
		tmp := int(word[i]-'a')
		wm[tmp] += 1
	}
	for k, v := range wm {
		val, ok := letters[k]
		if !ok {
			return false
		}
		if val < v {
			return false
		}
		letters[k] = val - v
	}
	return true
}
