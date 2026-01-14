package main

const ALPHABET_SIZE int = 26

type Node struct {
	value       rune
	isEndOfWord bool
	children    [ALPHABET_SIZE]*Node
}

func NewNode(val rune, isEndOfWord bool) *Node {
	return &Node{
		value:       val,
		isEndOfWord: isEndOfWord,
	}
}

type Trie struct {
	root *Node
}

func NewTrie() Trie {
	return Trie{
		root: &Node{},
	}
}

func (t *Trie) Insert(word string) {
	current := t.root

	for _, val := range word {
		// get the index of character and check if it is the end of the word
		index := val - 'a'
		if current.children[index] == nil {
			current.children[index] = NewNode(val, false)
		}
		current = current.children[index]
	}

	current.isEndOfWord = true
}
