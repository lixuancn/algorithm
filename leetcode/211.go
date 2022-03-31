package main

import "fmt"

//211. 添加与搜索单词 - 数据结构设计

type Node struct {
	char      rune
	childList map[rune]*Node
	isEnd     bool
}

func NewNode(char rune) *Node {
	return &Node{char: char, childList: map[rune]*Node{}}
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	root := NewNode('0')
	return WordDictionary{root: root}
}

func (this *WordDictionary) AddWord(word string) {
	parent := this.root
	for _, c := range word {
		if _, ok := parent.childList[c]; !ok {
			node := NewNode(c)
			parent.childList[c] = node
		}
		parent = parent.childList[c]
	}
	parent.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.SearchDFS(word, 0, this.root)
}

func (this *WordDictionary) SearchDFS(word string, index int, parent *Node) bool {
	if index == len(word) {
		return parent.isEnd
	}
	char := rune(word[index])
	if char != '.' {
		if _, ok := parent.childList[char]; !ok {
			return false
		}
		return this.SearchDFS(word, index+1, parent.childList[char])
	} else {
		for _, child := range parent.childList {
			if this.SearchDFS(word, index+1, child) {
				return true
			}
		}
	}
	return false
}

func main() {
	WordDictionary := Constructor()
	WordDictionary.AddWord("bad")
	WordDictionary.AddWord("dad")
	WordDictionary.AddWord("mad")
	fmt.Println(WordDictionary.Search("pad")) // 返回 False
	fmt.Println(WordDictionary.Search("bad")) // 返回 True
	fmt.Println(WordDictionary.Search(".ad")) // 返回 True
	fmt.Println(WordDictionary.Search("b..")) // 返回 True
}
