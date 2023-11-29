package leetcode

type WordDictionary struct {
	isWord   bool
	children map[rune]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{
		isWord:   false,
		children: make(map[rune]*WordDictionary),
	}
}

func (t *WordDictionary) AddWord(word string) {
	parent := t
	for _, c := range word {
		if child, ok := parent.children[c]; ok {
			parent = child
		} else {
			newChild := WordDictionary{children: make(map[rune]*WordDictionary)}
			parent.children[c] = &newChild
			parent = &newChild
		}
	}
	parent.isWord = true
}

func (t *WordDictionary) Search(word string) bool {
	parent := t
	for i, ch := range word {
		if rune(ch) == '.' {
			isMatched := false
			for _, v := range parent.children {
				if v.Search(word[i+1:]) {
					isMatched = true
				}
			}
			return isMatched
		} else if _, ok := parent.children[rune(ch)]; !ok {
			return false
		}
		parent = parent.children[rune(ch)]
	}
	return len(parent.children) == 0 || parent.isWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
