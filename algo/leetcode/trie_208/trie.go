package trie_208

// 实现一个字典树(字典树的边，存放具体元素)
// 用空间换时间
/*
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
*/

type Trie struct {
	value    string
	children map[string]*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{
		value:    "",
		children: make(map[string]*Trie),
		isWord:   false, //是否为一个完整的word，false表示：只是个前缀
	}
}

func (t *Trie) Insert(word string) {
	root := t
	for _, v := range word {
		n := root.match(string(v))
		if n != nil {
			root = n
		} else {
			newNode := &Trie{value: string(v), children: make(map[string]*Trie)}
			root.children[string(v)] = newNode
			root = root.children[string(v)]
		}
	}
	root.isWord = true
}

func (t *Trie) Search(word string) bool {
	if n := t.match(word); n != nil && n.isWord {
		return true
	} else {
		return false
	}
}

func (t *Trie) StartsWith(prefix string) bool {
	if n := t.match(prefix); n != nil {
		return true
	} else {
		return false
	}
}

func (t *Trie) match(word string) *Trie {
	root := t
	for _, v := range word {
		if _, ok := root.children[string(v)]; ok {
			root = root.children[string(v)]
		} else {
			return nil
		}
	}
	return root
}
