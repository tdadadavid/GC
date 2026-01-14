package trie

type NodeH struct {
	value       rune
	children    map[rune]*NodeH
	isEndOfWord bool
}

// trie implemented using a HashMap
type TrieH struct {
	root *NodeH
}

func NewNodeH(val rune, isEndOfWord bool) *NodeH {
	return &NodeH{
		value:       val,
		children:    make(map[rune]*NodeH, 0),
		isEndOfWord: isEndOfWord,
	}
}

func NewTrieH() TrieH {
	return TrieH{
		root: &NodeH{
			value:       '\x00',
			children:    map[rune]*NodeH{},
			isEndOfWord: false,
		},
	}
}

func (t *TrieH) InsertH(word string) {
	current := t.root

	for idx, val := range word {
		if current.children[val] == nil {
			isLastChar := idx == len(word)-1
			current.children[val] = NewNodeH(val, isLastChar)
		}
		current = current.children[val]
	}
}

func (t *TrieH) Contains(word string) (contains bool) {
	if len(word) == 0 {
		return contains
	}

	current := t.root
	for idx, val := range word {
		current = current.children[val]
		if current == nil {
			return contains
		}

		lastChar := idx == len(word)-1

		if val == current.value && lastChar && current.isEndOfWord {
			contains = true
			return contains
		}

	}

	return contains
}
