package trie

import (
	"testing"
)

// Test only focuses on HashMap implementation
func TestTrieH_Insert(t *testing.T) {
	tests := map[string]struct {
		inputs        []string
		expectedShape map[rune]*NodeH
	}{
		"Insert into trie": {
			inputs: []string{"boy", "boat", "blomblom"},
			expectedShape: map[rune]*NodeH{
				'\x00': {
					value:       '\x00',
					isEndOfWord: false,
					children: map[rune]*NodeH{
						'b': {
							value:       'b',
							isEndOfWord: false,
							children: map[rune]*NodeH{
								'o': {
									value:       'o',
									isEndOfWord: false,
									children: map[rune]*NodeH{
										'y': {
											value:       'y',
											isEndOfWord: true,
										},
										'a': {
											value:       'a',
											isEndOfWord: false,
											children: map[rune]*NodeH{
												't': {
													value:       't',
													isEndOfWord: true,
												},
											},
										},
									},
								},
								'l': {
									value:       'l',
									isEndOfWord: false,
									children: map[rune]*NodeH{
										'o': {
											value:       'o',
											isEndOfWord: false,
											children: map[rune]*NodeH{
												'm': {
													value:       'm',
													isEndOfWord: false,
													children: map[rune]*NodeH{
														'b': {
															value:       'b',
															isEndOfWord: false,
															children: map[rune]*NodeH{
																'l': {
																	value:       'l',
																	isEndOfWord: false,
																	children: map[rune]*NodeH{
																		'o': {
																			value:       'o',
																			isEndOfWord: false,
																			children: map[rune]*NodeH{
																				'm': {
																					value:       'm',
																					isEndOfWord: true,
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			trie := NewTrieH()

			for _, word := range test.inputs {
				trie.InsertH(word)
			}

			if compare(trie.root.children, test.expectedShape) {
				t.Fatalf("expected trie to be %+v got=%+v", test.expectedShape, trie)
			}
		})
	}
}

func TestTrieH_Contains(t *testing.T) {
	tests := map[string]struct {
		inputs        []string
		searchTerm    string
		shouldBeFound bool
	}{
		"Trie - Contains boy": {
			inputs:        []string{"boy", "boat", "blomblom"},
			searchTerm:    "boy",
			shouldBeFound: true,
		},
		"Trie - Contains bo": {
			inputs:        []string{"boy", "boat", "blomblom"},
			searchTerm:    "bo",
			shouldBeFound: false,
		},
		"Trie - Contains Empty string": {
			inputs:        []string{"boy", "boat", "blomblom"},
			searchTerm:    "",
			shouldBeFound: false,
		},
		"Trie - Contains blossom": {
			inputs:        []string{"boy", "boat", "blomblom"},
			searchTerm:    "blossom",
			shouldBeFound: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			trie := NewTrieH()

			for _, word := range test.inputs {
				trie.InsertH(word)
			}

			found := trie.Contains(test.searchTerm)
			if !test.shouldBeFound && found {
				t.Fatalf("expected word %s not be found got=%t", test.searchTerm, found)
			}
		})
	}
}

func compare(got, expected map[rune]*NodeH) bool {
	// nil vs nil is equal
	if got == nil && expected == nil {
		return true
	}

	// one nil, one not nil
	if got == nil || expected == nil {
		return false
	}

	// must have same number of keys
	if len(got) != len(expected) {
		return false
	}

	// compare each rune key and recurse
	for ch, gotNode := range got {
		expNode, ok := expected[ch]
		if !ok {
			return false // missing key
		}

		if !compareNode(gotNode, expNode) {
			return false
		}
	}

	return true
}

func compareNode(got, expected *NodeH) bool {
	// nil vs nil is equal
	if got == nil && expected == nil {
		return true
	}
	if got == nil || expected == nil {
		return false
	}

	// compare node fields (add/remove based on what you care about)
	if got.value != expected.value {
		return false
	}
	if got.isEndOfWord != expected.isEndOfWord {
		return false
	}

	// recurse into children
	return compare(got.children, expected.children)
}
