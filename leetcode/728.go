package main

func selfDividingNumbers(left int, right int) []int {
	ret := make([]int, 0)
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func isSelfDividing(target int) bool {
	num := target
	for num > 0 {
		i := num % 10
		if i == 0 || target%i != 0 {
			return false
		}
		num = num / 10
	}
	return true
}
