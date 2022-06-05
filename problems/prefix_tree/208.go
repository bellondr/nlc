package main

type Trie struct {
	isLeaf bool
	children [26]*Trie
}


func Constructor() Trie {
	return Trie{}
}


func (this *Trie) Insert(word string)  {
	for i := range word {
		if this.children[word[i] - 'a'] == nil {
			this.children[word[i] - 'a'] = &Trie{}
		}
		this = this.children[word[i] - 'a']
	}
	this.isLeaf = true
}


func (this *Trie) Search(word string) bool {
	for i := range word {
		if this.children[word[i] - 'a'] == nil {
			return false
		}
		this = this.children[word[i] - 'a']
	}
	return this.isLeaf
}


func (this *Trie) StartsWith(prefix string) bool {
	word := prefix
	for i := range word {
		if this.children[word[i] - 'a'] == nil {
			return false
		}
		this = this.children[word[i] - 'a']
	}
	return true
}

