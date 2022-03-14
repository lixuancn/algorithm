package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] > g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	gLen, sLen := len(g), len(s)
	sIndex := 0
	ret := 0
	for i := 0; i < gLen; i++ {
		if s[sIndex] >= g[i] && sIndex < sLen {
			sIndex++
			ret++
		}
	}
	return ret
}

func main() {
	//fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	//fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
