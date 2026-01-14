package main

import "encoding/json"

type jsonNode struct {
	Value       string               `json:"value"`
	IsEndOfWord bool                 `json:"isEndOfWord"`
	Children    map[string]*jsonNode `json:"children,omitempty"`
}

func (t Trie) String() string {
	if t.root == nil {
		return "{}"
	}

	root := toJSONNode(t.root)

	b, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(b)
}

func toJSONNode(n *Node) *jsonNode {
	if n == nil {
		return nil
	}

	out := &jsonNode{
		Value:       runeToString(n.value),
		IsEndOfWord: n.isEndOfWord,
		Children:    make(map[string]*jsonNode),
	}

	for i, child := range n.children {
		if child == nil {
			continue
		}

		// Convert index -> letter key ("a"..."z")
		key := string(rune('a' + i))
		out.Children[key] = toJSONNode(child)
	}

	// If no children, omit it in JSON
	if len(out.Children) == 0 {
		out.Children = nil
	}

	return out
}

type jsonNodeH struct {
	Value       string                `json:"value"`
	IsEndOfWord bool                  `json:"isEndOfWord"`
	Children    map[string]*jsonNodeH `json:"children,omitempty"`
}

func (t TrieH) String() string {
	if t.root == nil {
		return "{}"
	}

	root := toJSONNodeH(t.root)

	b, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(b)
}

func toJSONNodeH(n *NodeH) *jsonNodeH {
	if n == nil {
		return nil
	}

	out := &jsonNodeH{
		Value:       runeToString(n.value),
		IsEndOfWord: n.isEndOfWord,
	}

	if len(n.children) > 0 {
		out.Children = make(map[string]*jsonNodeH, len(n.children))

		for r, child := range n.children {
			out.Children[string(r)] = toJSONNodeH(child)
		}
	}

	return out
}

func runeToString(r rune) string {
	// root node typically has value = 0
	if r == 0 {
		return ""
	}
	return string(r)
}
