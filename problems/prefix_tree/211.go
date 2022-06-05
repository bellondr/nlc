package main

type WordDictionary struct {
	isLeaf bool
	children [26]*WordDictionary
}

/** Initialize your data structure here. */
func WConstructor() WordDictionary {
	return WordDictionary{}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	for i := range word {
		tmp := int(word[i] - 'a')
		if this.children[tmp] == nil {
			this.children[tmp] = &WordDictionary{}
		}
		this = this.children[tmp]
	}
	this.isLeaf = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	for i := range word {
		if word[i] == '.' {
			for _, v := range this.children {
				if v != nil && v.Search(word[i+1:]){
					return true
				}
			}
			return false
		}else {
			tmp := int(word[i] - 'a')
			if this.children[tmp] == nil {
				return false
			}
			this = this.children[tmp]
		}
	}
	return this.isLeaf
}
