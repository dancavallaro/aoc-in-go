package trie

type Trie struct {
	root *node
}

type node struct {
	value    rune
	isWord   bool
	children [26]*node
}

func newNode(v rune) *node {
	return &node{value: v, isWord: false}
}

func NewTrie() Trie {
	return Trie{root: newNode(0)}
}

func (trie Trie) Insert(word string) {
	curNode := trie.root
	for _, c := range []rune(word) {
		charNum := int(c - 'a')
		if curNode.children[charNum] == nil {
			curNode.children[charNum] = newNode(c)
		}
		curNode = curNode.children[charNum]
	}
	curNode.isWord = true
}

func (trie Trie) Contains(s string) (prefix bool, word bool) {
	curNode := trie.root
	for _, c := range []rune(s) {
		charNum := int(c - 'a')
		if curNode.children[charNum] == nil {
			prefix, word = false, false
			return
		}
		curNode = curNode.children[charNum]
	}
	prefix, word = true, curNode.isWord
	return
}

func (trie Trie) ContainsPrefix(s string) bool {
	prefix, _ := trie.Contains(s)
	return prefix
}

func (trie Trie) ContainsWord(s string) bool {
	_, word := trie.Contains(s)
	return word
}
