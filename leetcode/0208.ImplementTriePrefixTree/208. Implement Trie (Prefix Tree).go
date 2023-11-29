package leetcode

type Trie struct {
	isWord   bool
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		isWord:   false,
		children: make(map[rune]*Trie),
	}
}

func (t *Trie) Insert(word string) {
	parent := t
	for _, c := range word {
		if child, ok := parent.children[c]; ok {
			parent = child
		} else {
			newChild := &Trie{
				isWord:   false,
				children: make(map[rune]*Trie),
			}
			parent.children[c] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

func (t *Trie) Search(word string) bool {
	parent := t
	for _, c := range word {
		if child, ok := parent.children[c]; ok {
			parent = child
		} else {
			return false
		}
	}
	return parent.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	parent := t
	for _, c := range prefix {
		if child, ok := parent.children[c]; ok {
			parent = child
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
