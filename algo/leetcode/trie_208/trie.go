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

}

func (t *Trie) Insert(word string) {

}

func (t *Trie) Search(word string) bool {

}

func (t *Trie) StartsWith(prefix string) bool {

}
