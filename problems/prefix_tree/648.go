package main

import "strings"

type TireTree struct {
	children [26]*TireTree
	isLeaf bool
}
func replaceWords(dictionary []string, sentence string) string {
	tireTree := &TireTree{}
	for i := range dictionary {
		buildTree(dictionary[i], tireTree)
	}
	words := strings.Split(sentence, " ")
	for i := range words {
		tmp := rootWord(words[i], tireTree)
		words[i] = tmp
	}
	return strings.Join(words, " ")
}
func buildTree(word string, tree *TireTree) {
	this := tree
	for i := range word {
		if this.children[word[i] - 'a'] == nil {
			this.children[word[i] - 'a'] = &TireTree{}
		}
		this = this.children[word[i] - 'a']
	}
	this.isLeaf = true
}

func rootWord(word string, tree *TireTree) string {
	t := tree
	for i := range word {
		if t.children[word[i] - 'a'] == nil {
			return word
		}
		t = t.children[word[i] - 'a']
		if t.isLeaf {
			return word[:i+1]
		}
	}
	return word
}