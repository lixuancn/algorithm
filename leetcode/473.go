package main

import "fmt"

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}
	sum := 0
	for _, l := range matchsticks {
		sum += l
	}
	if sum%4 != 0 {
		return false
	}
	sum /= 4
	used := make([]bool, len(matchsticks))
	var backtracking func(matchsticks []int, target, bucket, k, start int, used []bool) bool
	backtracking = func(matchsticks []int, target, bucket, k, start int, used []bool) bool {
		if k == 0 {
			return true
		}
		if bucket == target {
			return backtracking(matchsticks, target, 0, k-1, 0, used)
		}
		for i := start; i < len(matchsticks); i++ {
			if used[i] {
				continue
			}
			if bucket+matchsticks[i] > target {
				continue
			}
			bucket += matchsticks[i]
			used[i] = true
			if backtracking(matchsticks, target, bucket, k, i+1, used) {
				return true
			}
			bucket -= matchsticks[i]
			used[i] = false
		}
		return false
	}
	return backtracking(matchsticks, sum, 0, 4, 0, used)
}

func main() {
	fmt.Println(makesquare([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 4, 3, 2, 1}))
	//fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
}
