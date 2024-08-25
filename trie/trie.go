package trie

// A trie is a prefix tree data structure
type Trie struct {
	children [26]*Trie
	// it have a node of 26 letters a-z and the node children are a-z also
	// Each element is a pointer to another Trie node, allowing the structure to branch out
	endOfWord bool
	// This boolean flag indicates whether the current node represents the end of a valid word.
	// It's crucial for distinguishing between complete words and prefixes.
	// For example, if you insert "app" and "apple", the node after "p" in "app" will have endOfWord set to true, while the same node in the path of "apple" will not
}

func Constructor() Trie {
	return Trie{}
}

func (d *Trie) Insert(word string) {
	for _, r := range word {
		i := r - 'a'
		if d.children[i] == nil {
			d.children[i] = &Trie{}
		}
		d = d.children[i]
	}
	d.endOfWord = true
}

func (d *Trie) Search(word string) bool {
	for _, r := range word {
		i := r - 'a'
		if d.children[i] == nil {
			return false
		}
		d = d.children[i]
	}
	return d.endOfWord
}

func (d *Trie) StartsWith(prefix string) bool {
	for _, r := range prefix {
		i := r - 'a'
		if d.children[i] == nil {
			return false
		}
		d = d.children[i]
	}
	return true
}

// WorldDictionary: Design a data structure that supports adding new words and finding if a string matches any previously added string.
type WorldDictionary struct {
	nodes        map[rune]*WorldDictionary
	endOfTheWord bool
}

// nodes is a map where each key is a rune (character) and the value is a pointer to another WorldDictionary. This allows for efficient character-by-character navigation.
// endOfTheWord is a boolean flag indicating if the current node represents the end of a valid word
func Init() WorldDictionary {
	return WorldDictionary{make(map[rune]*WorldDictionary), false}
}

func (this *WorldDictionary) AddWord(word string) {
	for _, r := range word {
		if _, found := this.nodes[r]; !found {
			node := Init()
			this.nodes[r] = &node
		}
		this = this.nodes[r]
	}
	this.endOfTheWord = true
}

// Iterates through each character of the word.
// If a character doesn't exist in the current node's map, it creates a new node for that character. Moves to the next node.
// After processing all characters, marks the last node as the end of a word.
func (this *WorldDictionary) Search(word string) bool {
	for i, r := range word {
		if _, found := this.nodes[r]; !found {
			if r != '.' {
				return false
			}
			for _, node := range this.nodes {
				if node.Search(word[i+1:]) {
					return true
				}
			}
			return false

		}
		this = this.nodes[r]
	}
	return this.endOfTheWord
}

// Iterates through each character of the search word.
// If a character is not found and it's not a '.', returns false.
// If the character is '.', it recursively searches all child nodes with the remaining substring.
// If all characters are processed, returns true only if it's the end of a word.
