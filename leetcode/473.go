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

func makesquare_practice(matchsticks []int) bool {
	targetLength := 0
	for _, v := range matchsticks {
		targetLength += v
	}
	if targetLength%4 != 0 {
		return false
	}
	targetLength /= 4
	used := make([]bool, len(matchsticks))
	var backtracking func(start, sideLength, sideCount int) bool
	backtracking = func(start, sideLength, sideCount int) bool {
		//四条边已经凑够
		if sideCount == 4 {
			return true
		}
		//一条边长度符合要求，开始凑下一条边
		if sideLength == targetLength {
			return backtracking(0, 0, sideCount+1)
		}
		for i := start; i < len(matchsticks); i++ {
			if sideLength+matchsticks[i] > targetLength || used[i] {
				continue
			}
			used[i] = true
			if backtracking(i+1, sideLength+matchsticks[i], sideCount) {
				return true
			}
			used[i] = false
		}
		return false
	}
	return backtracking(0, 0, 0)
}

func main() {
	fmt.Println(makesquare_practice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 4, 3, 2, 1}))
	fmt.Println(makesquare_practice([]int{1, 1, 2, 2, 2})) //true
	fmt.Println(makesquare_practice([]int{3, 3, 3, 3, 4})) //false
}
