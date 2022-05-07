package main

import (
	"container/list"
	"fmt"
)

//433. 最小基因变化

//bfs
func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	bankMap := map[string]struct{}{}
	for _, s := range bank {
		bankMap[s] = struct{}{}
	}
	if _, ok := bankMap[end]; !ok {
		return -1
	}
	queue := list.New()
	queue.PushBack(start)
	step := 0
	for queue.Len() > 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			cur := queue.Remove(queue.Front()).(string)
			for i, x := range cur {
				for _, y := range "ACGT" {
					if x == y {
						continue
					}
					newString := cur[:i] + string(y) + cur[i+1:]
					if _, ok := bankMap[newString]; !ok {
						continue
					}
					if newString == end {
						return step + 1
					}
					queue.PushBack(newString)
					delete(bankMap, newString)
				}
			}
		}
		step++
	}
	return -1
}

func main() {
	//fmt.Println(minMutation("AAAAAAAA", "AAAACAAA", []string{"AAACAAAA", "AAAAAAAA", "AAAACAAA"})) //1
	//fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"})) //1
	//fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"})) //2
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"})) //3
}
