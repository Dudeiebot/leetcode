package trie

type Trie struct {
	children [26]*Trie
	// it have a node of 26 letters a-z and the node children are a-z also
	endOfWord bool
	// we need end of the word because it practically returns true upon the end od the word
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
