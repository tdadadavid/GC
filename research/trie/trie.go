package main

import (
	"fmt"
)

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

		node := current.children[index]
		if node == nil {
			current.children[index] = NewNode(val, false)
		}
		current = current.children[index]
	}

	current.isEndOfWord = true
}

func main() {
	trieH := NewTrieH()
	trieH.InsertH("boy")
	trieH.InsertH("bomb")
	trieH.InsertH("boat")

	println()

	trie := NewTrie()
	trie.Insert("boy")
	trie.Insert("bomb")
	trie.Insert("boat")

	fmt.Printf("trieH %v\n\n", trieH.String())
	fmt.Printf("trie %v", trie.String())

}
