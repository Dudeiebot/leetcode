package trie

import "testing"

func TestTrie(t *testing.T) {
	words := []string{"apple", "app", "apricot", "banana"}
	trie := Constructor()

	for _, word := range words {
		trie.Insert(word)
	}

	testCase := []struct {
		word string
		want bool
	}{
		{"apple", true},
		{"app", true},
		{"apricot", true},
		{"banana", true},
		{"ap", false},
		{"application", false},
		{"orange", false},
	}

	for _, tc := range testCase {
		t.Run("Search for ", func(t *testing.T) {
			got := trie.Search(tc.word)
			if got != tc.want {
				t.Errorf("Search(%q) = %v, want %v", tc.word, got, tc.want)
			}
		})
	}

	testPrefix := []struct {
		prefix string
		want   bool
	}{
		{"app", true},
		{"apr", true},
		{"ban", true},
		{"bana", true},
		{"or", false},
		{"ba", true},
		{"c", false},
	}

	for _, tc := range testPrefix {
		t.Run("Stars with ", func(t *testing.T) {
			got := trie.StartsWith(tc.prefix)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// so a trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.
// so it intialize the trie Constructor
// and then inserts all the word inside it
// then we return true if the word is in the trie, nd false otherwise

func TestEndOfDictionary(t *testing.T) {
	dict := Init()
	testCases := []struct {
		addWords    []string
		searchWords map[string]bool
	}{
		{
			addWords: []string{"bad", "dad", "mad"},
			searchWords: map[string]bool{
				"pad": false,
				".ad": true,
				"b..": true,
				"b.d": true,
				"b.":  false,
				"...": true,
				"m.d": true,
				"mad": true,
				"ma.": true,
			},
		},
	}

	for _, tc := range testCases {
		for _, word := range tc.addWords {
			dict.AddWord(word)
		}

		for word, expected := range tc.searchWords {
			t.Run("Search_"+word, func(t *testing.T) {
				result := dict.Search(word)
				if result != expected {
					t.Errorf("Search(%q) = %v, want %v", word, result, expected)
				}
			})
		}
	}
}
