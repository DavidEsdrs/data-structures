package trie

type Node struct {
	Value       string
	IsEndOfWord bool
	Children    map[byte]*Node
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	root := NewNode()
	trie := Trie{root}
	return &trie
}

func NewNode() *Node {
	node := Node{IsEndOfWord: false, Children: map[byte]*Node{}}
	return &node
}

func (t *Trie) Insert(str string) {
	length := len(str)
	pCrawl := t.Root
	for level := 0; level < length; level++ {
		key := str[level]
		if _, found := pCrawl.Children[key]; !found {
			if pCrawl.Children == nil {
				pCrawl.Children = map[byte]*Node{}
			}
			pCrawl.Children[key] = &Node{IsEndOfWord: false}
		}
		pCrawl = pCrawl.Children[key]
	}
	pCrawl.IsEndOfWord = true
}

func (t *Trie) Search(str string) bool {
	length := len(str)
	pCrawl := t.Root
	for level := 0; level < length; level++ {
		key := str[level]
		if _, found := pCrawl.Children[key]; !found {
			return false
		}
		pCrawl = pCrawl.Children[key]
	}
	return pCrawl.IsEndOfWord
}
