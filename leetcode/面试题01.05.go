package main

import "fmt"

//面试题 01.05. 一次编辑

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m-n < -1 || m-n > 1 {
		return false
	}
	if first == second {
		return true
	}
	for i, j := 0, 0; i < m && j < n; i, j = i+1, j+1 {
		if first[i] != second[j] {
			if m == n {
				if first[i+1:] == second[j+1:] {
					return true
				}
				return false
			} else if m > n {
				if first[i+1:] == second[j:] {
					return true
				}
				return false
			} else {
				if first[i:] == second[j+1:] {
					return true
				}
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(oneEditAway("", "a"))
	fmt.Println(oneEditAway("a", "ab"))
}
