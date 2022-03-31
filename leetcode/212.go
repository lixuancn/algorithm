package main

import "fmt"

//212. 单词搜索 II

var strList map[string]struct{}
var retList []string
var resultList map[string]struct{}
var usedList [][]bool

func findWords(board [][]byte, words []string) []string {
	//return findWords_recursion(board, words) //回溯，但是剪枝不够充分
	return findWords_trie(board, words) //用字典树来剪纸
}

//用字典树
var trie Trie

func findWords_trie(board [][]byte, words []string) []string {
	trie = Constructor()
	for _, word := range words {
		trie.Insert(word)
	}
	resultList = make(map[string]struct{}, 0)
	usedList = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		usedList[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if trie.StartsWith(string(board[i][j])) {
				usedList[i][j] = true
				backtracking_trie(board, string(board[i][j]), i, j)
				usedList[i][j] = false
			}
		}
	}
	keys := make([]string, len(resultList))
	i := 0
	for k := range resultList {
		keys[i] = k
		i++
	}
	return keys
}

func backtracking_trie(board [][]byte, str string, i, j int) {
	//题意，单个单词长度不超过10
	if len(str) > 10 {
		return
	}
	if trie.Search(str) {
		resultList[str] = struct{}{}
	}
	for _, pos := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		newI, newJ := i+pos[0], j+pos[1]
		if newI < 0 || newI >= len(board) || newJ < 0 || newJ == len(board[0]) {
			continue
		}
		if usedList[newI][newJ] == true {
			continue
		}
		usedList[newI][newJ] = true
		backtracking_trie(board, str+string(board[newI][newJ]), newI, newJ)
		usedList[newI][newJ] = false
	}
}

//回溯，但是剪枝不够充分
func findWords_recursion(board [][]byte, words []string) []string {
	retList = make([]string, 0)
	strList = map[string]struct{}{}
	usedList = make([][]bool, len(board))
	for _, word := range words {
		strList[word] = struct{}{}
	}
	for i := 0; i < len(board); i++ {
		usedList[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			usedList[i][j] = true
			backtracking(board, string(board[i][j]), i, j)
			usedList[i][j] = false
		}
	}
	return retList
}

func backtracking(board [][]byte, str string, i, j int) {
	//题意，单个单词长度不超过10
	if len(str) > 10 {
		return
	}
	if _, ok := strList[str]; ok {
		retList = append(retList, str)
		delete(strList, str)
	}
	for _, pos := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		newI, newJ := i+pos[0], j+pos[1]
		if newI < 0 || newI >= len(board) || newJ < 0 || newJ == len(board[0]) {
			continue
		}
		if usedList[newI][newJ] == true {
			continue
		}
		usedList[newI][newJ] = true
		backtracking(board, str+string(board[newI][newJ]), newI, newJ)
		usedList[newI][newJ] = false
	}
}

type Node struct {
	char      rune
	childList map[rune]*Node
	isEnd     bool
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
	parent.isEnd = true
}

func (this *Trie) Search(word string) bool {
	parent := this.root
	for _, c := range word {
		if _, ok := parent.childList[c]; !ok {
			return false
		}
		parent = parent.childList[c]
	}
	if parent.isEnd {
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
	//fmt.Println(findWords([][]byte{{'a'}}, []string{"a"}))
	fmt.Println(findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}))
}
