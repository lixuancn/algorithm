package main

import "fmt"

type Node struct {
	char      rune
	childList map[rune]*Node
}

func NewNode(char rune) *Node {
	return &Node{char: char, childList: map[rune]*Node{}}
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	root := NewNode('0')
	return Trie{root: root}
}

func (this *Trie) Insert(word string) {
	parent := this.root
	for _, c := range word {
		if _, ok := parent.childList[c]; !ok {
			node := NewNode(c)
			parent.childList[c] = node
		}
		parent = parent.childList[c]
	}
	parent.childList[0] = nil
}

func (this *Trie) Search(word string) bool {
	parent := this.root
	for _, c := range word {
		if _, ok := parent.childList[c]; !ok {
			return false
		}
		parent = parent.childList[c]
	}
	if _, ok := parent.childList[0]; ok {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this.root
	for _, c := range prefix {
		if _, ok := parent.childList[c]; !ok {
			return false
		}
		parent = parent.childList[c]
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // 返回 True
	fmt.Println(trie.Search("app"))     // 返回 False
	fmt.Println(trie.StartsWith("app")) // 返回 True
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // 返回 True
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
