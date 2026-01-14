package main

import "fmt"

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

// Root -> Children
func (t *TrieH) PreOrder() (result []string) {
	if t.root == nil {
		return result
	}

	preOrder(t.root, &result)
	return result
}

// Children -> Root
func (t *TrieH) PostOrder() (result []string) {
	if t.root == nil {
		return result
	}

	postOrder(t.root, &result)
	return result
}

func postOrder(node *NodeH, result *[]string) {
	for _, child := range node.children { // Children
		postOrder(child, result)
	}

	*result = append(*result, string(node.value))
}

func preOrder(node *NodeH, result *[]string) {
	*result = append(*result, string(node.value)) // Root

	for _, child := range node.children { // Children
		preOrder(child, result)
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

func main() {
	trieH := NewTrieH()
	trieH.InsertH("care")
	// trieH.InsertH("bomb")
	// trieH.InsertH("boat")
	fmt.Printf("trieH contains boy: %t\n", trieH.Contains("boy"))
	fmt.Printf("trieH contains bomb: %t\n", trieH.Contains("bomb"))
	fmt.Printf("trieH contains bomber: %t\n", trieH.Contains("bomber"))
	fmt.Printf("trieH contains bo: %t\n", trieH.Contains("bo"))
	fmt.Printf("trieH contains `''`: %t\n", trieH.Contains(""))
	fmt.Printf("trieH contains boat: %t\n", trieH.Contains("boat"))
	fmt.Printf("trieH contains ``: %t\n", trieH.Contains(``))
	// fmt.Printf("trieH %v\n", trieH.String())

	fmt.Printf("result(postorder) %v\n", trieH.PostOrder())
	fmt.Printf("result(preorder) %v\n", trieH.PreOrder())
}
