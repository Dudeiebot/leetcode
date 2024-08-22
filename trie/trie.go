package trie

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
