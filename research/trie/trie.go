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
	fmt.Printf("trieH contains boy: %t\n", trieH.Contains("boy"))
	fmt.Printf("trieH contains bomb: %t\n", trieH.Contains("bomb"))
	fmt.Printf("trieH contains bomber: %t\n", trieH.Contains("bomber"))
	fmt.Printf("trieH contains bo: %t\n", trieH.Contains("bo"))
	fmt.Printf("trieH contains `''`: %t\n", trieH.Contains(""))
	fmt.Printf("trieH contains ``: %t\n", trieH.Contains(``))

	println()

	trie := NewTrie()
	trie.Insert("boy")
	trie.Insert("bomb")
	trie.Insert("boat")

	// fmt.Printf("trieH contains boy: %t\n", trie.Contains("boy"))
	// fmt.Printf("trieH contains bo: %t\n", trie.Contains("bo"))

	// fmt.Printf("trieH %v\n\n", trieH.String())
	// fmt.Printf("trie %v", trie.String())

}
