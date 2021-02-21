package main

type Tire struct {
	children [26]*Tire
	isLeaf bool
}

type WordDictionary struct {
	root *Tire
}


/** Initialize your data structure here. */
func WDConstructor() WordDictionary {
	return WordDictionary {
		root: &Tire{},
	}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	root := this.root
	addword(root, word)

}

func addword(root *Tire, word string) {
	if len(word) == 0 {
		root.isLeaf = true
		return
	}
	for i, w := range word {
		if w != '.' {
			index := w - 'a'
			if root.children[index] == nil {
				root.children[index] = &Tire{}
			}
			root = root.children[index]
		} else {
			for ni := 0; ni < 26; ni++ {
				if root.children[ni] == nil {
					root.children[ni] = &Tire{}
				}
				addword(root.children[ni], word[i+1:])
			}
		}
	}
	root.isLeaf = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return searchWorld(this.root, word)
}

func searchWorld(root *Tire, word string) bool {
	if len(word) == 0 {
		return root.isLeaf
	}
	for i, w := range word {
		if w != '.' {
			index := w - 'a'
			if root.children[index] == nil {
				return false
			}
			root = root.children[index]
		} else {
			for ni, c := range root.children {
				if c != nil {
					if searchWorld(root.children[ni], word[i+1:]) {
						return true
					}
				}
			}
			return false
		}
	}
	return root.isLeaf
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
