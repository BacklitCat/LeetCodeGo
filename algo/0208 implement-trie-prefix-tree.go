package algo

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
	parent *Trie
	kids   []*Trie // a-97, z-122
	cut    bool
}

func TrieConstructor() Trie {
	return Trie{
		parent: nil,
		kids:   make([]*Trie, 26),
	}
}

func (this *Trie) Insert(word string) {
	p := this
	for _, s := range word {
		if p.kids[s-97] != nil {
			p = p.kids[s-97]
			continue
		} else {
			node := &Trie{
				parent: p,
				kids:   make([]*Trie, 26),
			}
			p.kids[s-97] = node
			p = node
		}
	}
	p.cut = true
}

func (this *Trie) Search(word string) bool {
	p := this
	for _, s := range word {
		if p.kids[s-97] == nil {
			return false
		}
		p = p.kids[s-97]
	}

	return p.cut
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for _, s := range prefix {
		if p.kids[s-97] == nil {
			return false
		}
		p = p.kids[s-97]
	}
	return true
}
