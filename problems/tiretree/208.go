package main
type Trie struct {
	isLeaf bool
	children [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	for i := range word {
		if this.children[word[i]-'a'] == nil {
			this.children[word[i]-'a'] = &Trie{}
		}
		this =  this.children[word[i]-'a']
	}
	this.isLeaf = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for i := range word {
		if this.children[word[i] - 'a'] == nil {
			return false
		}
		this = this.children[word[i] - 'a']
	}
	return this.isLeaf
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for i := range prefix {
		if this.children[prefix[i] - 'a'] == nil {
			return false
		}
		this = this.children[prefix[i] - 'a']
	}
	return true
}
