package main

import (
	"fmt"
	"sort"
)

//721. 账户合并

//并查集
func accountsMerge(accounts [][]string) [][]string {
	emailToIndex := make(map[string]int)
	emailToName := make(map[string]string)
	for _, account := range accounts {
		for _, email := range account[1:] {
			if _, ok := emailToIndex[email]; !ok {
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = account[0]
			}
		}
	}

	parent := make([]int, len(emailToIndex))
	for i := range parent {
		parent[i] = i
	}
	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]]
		for _, email := range account[2:] {
			union(parent, firstIndex, emailToIndex[email]) //
		}
	}

	indexToEmail := map[int][]string{}
	for email, index := range emailToIndex {
		index = find(parent, index)
		indexToEmail[index] = append(indexToEmail[index], email)
	}

	ans := [][]string{}
	for _, emails := range indexToEmail {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return ans
}

func find(parent []int, index int) int {
	for parent[index] != index {
		parent[index] = parent[parent[index]]
		index = parent[index]
	}
	return index
}

func union(parent []int, index1, index2 int) {
	parent[find(parent, index1)] = find(parent, index2)
}

func main() {
	//fmt.Println(accountsMerge([][]string{
	//	{"John", "johnsmith@mail.com", "john00@mail.com"},
	//	{"John", "johnnybravo@mail.com"},
	//	{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
	//	{"Mary", "mary@mail.com"},
	//}))

	fmt.Println(accountsMerge([][]string{
		{"a", "1", "2"},
		{"a", "3"},
		{"a", "1", "4"},
		{"b", "5"},
	}))
}
