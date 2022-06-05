package main

type MapSum struct {
	children [26]*MapSum
	val int
}


func MapSumConstructor() MapSum {
	return MapSum{}
}


func (this *MapSum) Insert(key string, val int)  {
	for i := range key {
		if this.children[key[i] - 'a'] == nil {
			this.children[key[i] - 'a'] = &MapSum{}
		}
		this = this.children[key[i] - 'a']
	}
	this.val = val
}


func (this *MapSum) Sum(prefix string) int {
	tree := this
	for i := range prefix {
		if tree.children[prefix[i] - 'a'] == nil {
			return 0
		}
		tree = tree.children[prefix[i] - 'a']
	}
	return tree.sumAll()
}

func (this *MapSum) sumAll() int {
	if this == nil {
		return 0
	}
	res := this.val
	for i:= 0; i < 26; i++ {
		if this.children[i] != nil {
			tmp := this.children[i].sumAll()
			res += tmp
		}
	}
	return res
}