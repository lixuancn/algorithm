package main

//给一非空的单词列表，返回前 k 个出现次数最多的单词。
//返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	param := []string{"plpaboutit","jnoqzdute","sfvkdqf","mjc","nkpllqzjzp","foqqenbey","ssnanizsav","nkpllqzjzp","sfvkdqf","isnjmy","pnqsz","hhqpvvt","fvvdtpnzx","jkqonvenhx","cyxwlef","hhqpvvt","fvvdtpnzx","plpaboutit","sfvkdqf","mjc","fvvdtpnzx","bwumsj","foqqenbey","isnjmy","nkpllqzjzp","hhqpvvt","foqqenbey","fvvdtpnzx","bwumsj","hhqpvvt","fvvdtpnzx","jkqonvenhx","jnoqzdute","foqqenbey","jnoqzdute","foqqenbey","hhqpvvt","ssnanizsav","mjc","foqqenbey","bwumsj","ssnanizsav","fvvdtpnzx","nkpllqzjzp","jkqonvenhx","hhqpvvt","mjc","isnjmy","bwumsj","pnqsz","hhqpvvt","nkpllqzjzp","jnoqzdute","pnqsz","nkpllqzjzp","jnoqzdute","foqqenbey","nkpllqzjzp","hhqpvvt","fvvdtpnzx","plpaboutit","jnoqzdute","sfvkdqf","fvvdtpnzx","jkqonvenhx","jnoqzdute","nkpllqzjzp","jnoqzdute","fvvdtpnzx","jkqonvenhx","hhqpvvt","isnjmy","jkqonvenhx","ssnanizsav","jnoqzdute","jkqonvenhx","fvvdtpnzx","hhqpvvt","bwumsj","nkpllqzjzp","bwumsj","jkqonvenhx","jnoqzdute","pnqsz","foqqenbey","sfvkdqf","sfvkdqf"}
	r := topKFrequent(param, 1)
	fmt.Println(r)
}

//优先队列
func topKFrequent(words []string, k int) []string {
	countList := make(map[string]int)
	for _, word := range words{
		countList[word]++
	}
	pq := &priorityQueue{}
	for word, num := range countList{
		heap.Push(pq, item{word:word, num:num})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	ret := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ret[i] = heap.Pop(pq).(item).word
	}
	return ret
}

type item struct{
	word string
	num int
}
type priorityQueue []item
func (pq priorityQueue)Len()int { return len(pq) }
func (pq priorityQueue)Less(i, j int) bool {
	a, b := pq[i], pq[j]
	return a.num < b.num || a.num == b.num && a.word > b.word
}
func (pq priorityQueue)Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue)Push(v interface{}) {
	*pq = append(*pq, v.(item))
}
func (pq *priorityQueue)Pop() interface{} {
	a := *pq
	v := a[len(a)-1]
	*pq = a[:len(a)-1]
	return v
}

//hash表的方法
func topKFrequent_hash(words []string, k int) []string {
	countList := make(map[string]int)
	for _, word := range words{
		countList[word]++
	}
	wordList := make([]string, 0)
	for word := range countList{
		wordList = append(wordList, word)
	}
	//数字从大到小排序，字符串从小到大
	sort.Slice(wordList, func (i, j int) bool {
		return countList[wordList[i]] > countList[wordList[j]] ||
			countList[wordList[i]] == countList[wordList[j]] && wordList[i] < wordList[j]
	})
	return wordList[:k]
}
